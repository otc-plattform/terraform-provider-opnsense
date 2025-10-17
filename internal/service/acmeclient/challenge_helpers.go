package acmeclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mitchellh/mapstructure"
)

func challengeResponseToModel(resp *acmeclient.ValidationGetResponse) acmeclientChallengeResourceModel {
	validation := resp.Validation

	model := acmeclientChallengeResourceModel{
		Id:                       types.StringNull(),
		Enabled:                  types.BoolValue(tools.StringToBool(validation.Enabled)),
		Name:                     stringValueOrNull(validation.Name),
		Description:              stringValueOrNull(validation.Description),
		Method:                   stringValueFromOptionMap(validation.Method),
		HTTPService:              stringValueFromOptionMap(validation.HTTPService),
		HTTPOpnAutodiscovery:     types.BoolValue(tools.StringToBool(validation.HTTPOpnAutodiscovery)),
		HTTPOpnInterface:         stringValueFromOptionMap(validation.HTTPOpnInterface),
		HTTPOpnIPAddresses:       stringSliceToSet(validation.HTTPOpnIPAddresses),
		HTTPHAProxyInject:        types.BoolValue(tools.StringToBool(validation.HTTPHAProxyInject)),
		HTTPHAProxyFrontends:     stringSliceToSet(validation.HTTPHAProxyFrontends),
		TLSALPNService:           stringValueFromOptionMap(validation.TLSALPNService),
		TLSALPNAcmeAutodiscovery: types.BoolValue(tools.StringToBool(validation.TLSALPNAcmeAutodiscovery)),
		TLSALPNAcmeInterface:     stringValueFromOptionMap(validation.TLSALPNAcmeInterface),
		TLSALPNAcmeIPAddresses:   stringSliceToSet(validation.TLSALPNAcmeIPAddresses),
		DNSService:               stringValueFromOptionMap(validation.DNSService),
		DNSSleep:                 types.Int64Value(tools.StringToInt64(validation.DNSSleep)),
		Parameters:               stringMapToTypesMap(collectChallengeParameters(validation)),
	}

	return model
}

func challengeSearchRowToModel(challenge acmeclient.ChallengeSearchItem) acmeclientChallengeResourceModel {
	return acmeclientChallengeResourceModel{
		Id:          stringValueOrNull(challenge.UUID),
		Enabled:     types.BoolValue(tools.StringToBool(challenge.Enabled)),
		Name:        stringValueOrNull(challenge.Name),
		Description: stringValueOrNull(challenge.Description),
		Method:      stringValueOrNull(challenge.Method),
	}
}

func fetchChallengeModel(ctx context.Context, controller *acmeclient.Controller, id string) (acmeclientChallengeResourceModel, error) {
	response, err := controller.ACMEClientGetChallengeType(ctx, id)
	if err == nil {
		result := challengeResponseToModel(response)
		if result.Id.IsNull() || result.Id.ValueString() == "" {
			result.Id = types.StringValue(id)
		}
		return result, nil
	}

	var notFoundErr *errs.NotFoundError
	if errors.As(err, &notFoundErr) {
		return acmeclientChallengeResourceModel{}, err
	}

	if !isAccountArrayDecodeError(err) {
		return acmeclientChallengeResourceModel{}, err
	}

	searchResp, searchErr := controller.ACMEClientSearchChallengeType(ctx)
	if searchErr != nil {
		return acmeclientChallengeResourceModel{}, fmt.Errorf("fallback search failed: %w", searchErr)
	}

	for _, row := range searchResp.Rows {
		if row.UUID == id {
			result := challengeSearchRowToModel(row)
			if result.Id.IsNull() || result.Id.ValueString() == "" {
				result.Id = types.StringValue(id)
			}
			return result, nil
		}
	}

	return acmeclientChallengeResourceModel{}, errs.NewNotFoundError()
}

func (m *acmeclientChallengeResourceModel) toValidation(ctx context.Context) (acmeclient.Validation, error) {
	data := map[string]any{
		"enabled":                    boolToAPIString(m.Enabled),
		"name":                       stringToAPIValue(m.Name),
		"description":                stringToAPIValue(m.Description),
		"method":                     stringToAPIValue(m.Method),
		"http_service":               stringToAPIValue(m.HTTPService),
		"http_opn_autodiscovery":     boolToAPIString(m.HTTPOpnAutodiscovery),
		"http_opn_interface":         stringToAPIValue(m.HTTPOpnInterface),
		"http_opn_ipaddresses":       joinStringSet(ctx, m.HTTPOpnIPAddresses),
		"http_haproxyInject":         boolToAPIString(m.HTTPHAProxyInject),
		"http_haproxyFrontends":      joinStringSet(ctx, m.HTTPHAProxyFrontends),
		"tlsalpn_service":            stringToAPIValue(m.TLSALPNService),
		"tlsalpn_acme_autodiscovery": boolToAPIString(m.TLSALPNAcmeAutodiscovery),
		"tlsalpn_acme_interface":     stringToAPIValue(m.TLSALPNAcmeInterface),
		"tlsalpn_acme_ipaddresses":   joinStringSet(ctx, m.TLSALPNAcmeIPAddresses),
		"dns_service":                stringToAPIValue(m.DNSService),
		"dns_sleep":                  int64ToAPIString(m.DNSSleep),
	}

	if !m.Parameters.IsNull() && !m.Parameters.IsUnknown() {
		var params map[string]string
		if err := m.Parameters.ElementsAs(ctx, &params, false); err == nil {
			for k, v := range params {
				data[k] = v
			}
		}
	}

	var validation acmeclient.Validation
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &validation,
	})
	if err != nil {
		return acmeclient.Validation{}, fmt.Errorf("unable to create decoder for validation payload: %w", err)
	}

	if err := decoder.Decode(data); err != nil {
		return acmeclient.Validation{}, fmt.Errorf("unable to decode validation payload: %w", err)
	}

	return validation, nil
}

func stringSliceToSet(values []string) types.Set {
	if len(values) == 0 {
		return types.SetNull(types.StringType)
	}

	return tools.StringSliceToSet(values)
}

func joinStringSet(ctx context.Context, set types.Set) string {
	if set.IsNull() || set.IsUnknown() {
		return ""
	}

	list := tools.SetToStringSlice(set)
	return strings.Join(list, ",")
}

func stringMapToTypesMap(values map[string]string) types.Map {
	if len(values) == 0 {
		return types.MapNull(types.StringType)
	}

	attrValues := make(map[string]attr.Value, len(values))
	for k, v := range values {
		attrValues[k] = types.StringValue(v)
	}

	mapValue, _ := types.MapValue(types.StringType, attrValues)
	return mapValue
}

func collectChallengeParameters(validation acmeclient.ValidationGet) map[string]string {
	knownKeys := map[string]struct{}{
		"id":                         {},
		"enabled":                    {},
		"name":                       {},
		"description":                {},
		"method":                     {},
		"http_service":               {},
		"http_opn_autodiscovery":     {},
		"http_opn_interface":         {},
		"http_opn_ipaddresses":       {},
		"http_haproxyInject":         {},
		"http_haproxyFrontends":      {},
		"tlsalpn_service":            {},
		"tlsalpn_acme_autodiscovery": {},
		"tlsalpn_acme_interface":     {},
		"tlsalpn_acme_ipaddresses":   {},
		"dns_service":                {},
		"dns_sleep":                  {},
		"key":                        {},
		"statusCode":                 {},
		"statusLastUpdate":           {},
	}

	raw := map[string]any{}
	data, err := json.Marshal(validation)
	if err != nil {
		return map[string]string{}
	}
	if err := json.Unmarshal(data, &raw); err != nil {
		return map[string]string{}
	}

	params := map[string]string{}
	for key, value := range raw {
		if _, exists := knownKeys[key]; exists {
			continue
		}

		switch v := value.(type) {
		case string:
			if v == "" {
				continue
			}
			params[key] = v
		}
	}

	return params
}
