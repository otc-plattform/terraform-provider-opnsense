package acmeclient

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/mitchellh/mapstructure"
)

func certificateResponseToModel(resp *acmeclient.CertificateGetResponse) acmeclientCertificateResourceModel {
	cert := resp.Certificate

	return acmeclientCertificateResourceModel{
		Id:               types.StringNull(),
		Enabled:          types.BoolValue(tools.StringToBool(cert.Enabled)),
		Name:             stringValueOrNull(cert.Name),
		Description:      stringValueOrNull(cert.Description),
		AltNames:         stringSliceToSet(selectedOptionValues(cert.AltNames)),
		Account:          stringValueFromOptionMap(cert.Account),
		ValidationMethod: stringValueFromOptionMap(cert.ValidationMethod),
		AutoRenewal:      types.BoolValue(tools.StringToBool(cert.AutoRenewal)),
		RenewInterval:    types.Int64Value(tools.StringToInt64(cert.RenewInterval)),
		KeyLength:        stringValueFromOptionMap(cert.KeyLength),
		OCSP:             types.BoolValue(tools.StringToBool(cert.OCSP)),
		RestartActions:   stringSliceToSet(selectedOptionKeys(cert.RestartActions)),
		AliasMode:        stringValueFromOptionMap(cert.AliasMode),
		DomainAlias:      stringValueOrNull(cert.DomainAlias),
		ChallengeAlias:   stringValueOrNull(cert.ChallengeAlias),
		CertificateRefId: stringValueOrNull(cert.CertRefID),
		LastUpdate:       stringValueOrNull(cert.LastUpdate),
		StatusCode:       stringValueOrNull(cert.StatusCode),
		StatusLastUpdate: stringValueOrNull(cert.StatusLastUpdate),
	}
}

func certificateSearchRowToModel(cert acmeclient.CertificateSearchItem) acmeclientCertificateResourceModel {
	return acmeclientCertificateResourceModel{
		Id:               stringValueOrNull(cert.UUID),
		Enabled:          types.BoolValue(tools.StringToBool(cert.Enabled)),
		Name:             stringValueOrNull(cert.Name),
		Description:      stringValueOrNull(cert.Description),
		AltNames:         stringSliceToSet(strings.Split(cert.AltNames, ",")),
		StatusCode:       stringValueOrNull(cert.StatusCode),
		StatusLastUpdate: stringValueOrNull(cert.StatusLastUpdate),
		LastUpdate:       stringValueOrNull(cert.LastUpdate),
	}
}

func fetchCertificateModel(ctx context.Context, controller *acmeclient.Controller, id string) (acmeclientCertificateResourceModel, error) {
	resp, err := controller.ACMEClientGetCert(ctx, id)
	if err == nil {
		model := certificateResponseToModel(resp)
		if model.Id.IsNull() || model.Id.ValueString() == "" {
			model.Id = types.StringValue(id)
		}
		return model, nil
	}

	var notFoundErr *errs.NotFoundError
	if errors.As(err, &notFoundErr) {
		return acmeclientCertificateResourceModel{}, err
	}

	if !isAccountArrayDecodeError(err) {
		return acmeclientCertificateResourceModel{}, err
	}

	search, searchErr := controller.ACMEClientSearchCert(ctx)
	if searchErr != nil {
		return acmeclientCertificateResourceModel{}, fmt.Errorf("fallback search failed: %w", searchErr)
	}

	for _, row := range search.Rows {
		if row.UUID == id {
			model := certificateSearchRowToModel(row)
			if model.Id.IsNull() || model.Id.ValueString() == "" {
				model.Id = types.StringValue(id)
			}
			return model, nil
		}
	}

	return acmeclientCertificateResourceModel{}, errs.NewNotFoundError()
}

func (m *acmeclientCertificateResourceModel) toCertificate(ctx context.Context) (acmeclient.Certificate, error) {
	altNames := joinStringSet(ctx, m.AltNames)
	restartActions := joinStringSet(ctx, m.RestartActions)

	payload := map[string]any{
		"enabled":          boolToAPIString(m.Enabled),
		"name":             stringToAPIValue(m.Name),
		"description":      stringToAPIValue(m.Description),
		"altNames":         altNames,
		"account":          stringToAPIValue(m.Account),
		"validationMethod": stringToAPIValue(m.ValidationMethod),
		"autoRenewal":      boolToAPIString(m.AutoRenewal),
		"renewInterval":    int64ToAPIString(m.RenewInterval),
		"keyLength":        stringToAPIValue(m.KeyLength),
		"ocsp":             boolToAPIString(m.OCSP),
		"restartActions":   restartActions,
		"aliasmode":        stringToAPIValue(m.AliasMode),
		"domainalias":      stringToAPIValue(m.DomainAlias),
		"challengealias":   stringToAPIValue(m.ChallengeAlias),
	}

	var cert acmeclient.Certificate
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName: "json",
		Result:  &cert,
	})
	if err != nil {
		return acmeclient.Certificate{}, fmt.Errorf("unable to create decoder for certificate payload: %w", err)
	}

	if err := decoder.Decode(payload); err != nil {
		return acmeclient.Certificate{}, fmt.Errorf("unable to decode certificate payload: %w", err)
	}

	return cert, nil
}

func selectedOptionValues(options map[string]acmeclient.Option) []string {
	if len(options) == 0 {
		return []string{}
	}

	var entries []string
	for key, option := range options {
		if option.Selected != 1 {
			continue
		}

		value := option.Value
		if value == "" {
			value = key
		}

		entries = append(entries, value)
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i] < entries[j]
	})

	return entries
}

func selectedOptionKeys(options map[string]acmeclient.Option) []string {
	if len(options) == 0 {
		return []string{}
	}

	var keys []string
	for key, option := range options {
		if option.Selected == 1 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)
	return keys
}
