package acmeclient

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func accountResponseToModel(resp *acmeclient.AccountGetResponse) acmeclientAccountResourceModel {
	account := resp.Account

	return acmeclientAccountResourceModel{
		Id:               types.StringNull(),
		Enabled:          types.BoolValue(tools.StringToBool(account.Enabled)),
		Name:             stringValueOrNull(account.Name),
		Description:      stringValueOrNull(account.Description),
		Email:            stringValueOrNull(account.Email),
		CA:               stringValueFromOptionMap(account.CA),
		CustomCA:         stringValueOrNull(account.CustomCA),
		EABKID:           stringValueOrNull(account.EABKID),
		EABHMAC:          stringValueOrNull(account.EABHMAC),
		Key:              stringValueOrNull(account.Key),
		StatusCode:       stringValueOrNull(account.StatusCode),
		StatusLastUpdate: stringValueOrNull(account.StatusLastUpdate),
	}
}

func accountSearchRowToModel(account acmeclient.Account) acmeclientAccountResourceModel {
	return acmeclientAccountResourceModel{
		Id:               stringValueOrNull(account.UUID),
		Enabled:          types.BoolValue(tools.StringToBool(account.Enabled)),
		Name:             stringValueOrNull(account.Name),
		Description:      stringValueOrNull(account.Description),
		Email:            stringValueOrNull(account.Email),
		CA:               stringValueOrNull(account.CA),
		CustomCA:         stringValueOrNull(account.CustomCA),
		EABKID:           stringValueOrNull(account.EABKID),
		EABHMAC:          stringValueOrNull(account.EABHMAC),
		Key:              types.StringNull(),
		StatusCode:       types.StringNull(),
		StatusLastUpdate: types.StringNull(),
	}
}

func fetchAccountModel(ctx context.Context, controller *acmeclient.Controller, id string) (acmeclientAccountResourceModel, error) {
	accountResp, err := controller.ACMEClientGetAccount(ctx, id)
	if err == nil {
		model := accountResponseToModel(accountResp)
		if model.Id.IsNull() || model.Id.ValueString() == "" {
			model.Id = types.StringValue(id)
		}
		return model, nil
	}

	var notFoundErr *errs.NotFoundError
	if errors.As(err, &notFoundErr) {
		return acmeclientAccountResourceModel{}, err
	}

	if !isAccountArrayDecodeError(err) {
		return acmeclientAccountResourceModel{}, err
	}

	searchResp, searchErr := controller.ACMEClientSearchAccount(ctx)
	if searchErr != nil {
		return acmeclientAccountResourceModel{}, fmt.Errorf("fallback search failed: %w", searchErr)
	}

	for _, row := range searchResp.Rows {
		if row.UUID == id {
			model := accountSearchRowToModel(row)
			if model.Id.IsNull() || model.Id.ValueString() == "" {
				model.Id = types.StringValue(id)
			}
			return model, nil
		}
	}

	return acmeclientAccountResourceModel{}, errs.NewNotFoundError()
}

func isAccountArrayDecodeError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "cannot unmarshal array")
}

func stringValueOrNull(value string) types.String {
	return types.StringValue(value)
}

func (m *acmeclientAccountResourceModel) toAccountCreateRequest() acmeclient.AccountCreateRequest {
	return acmeclient.AccountCreateRequest{
		Enabled:     boolToAPIString(m.Enabled),
		Name:        stringToAPIValue(m.Name),
		Description: stringToAPIValue(m.Description),
		Email:       stringToAPIValue(m.Email),
		CA:          stringToAPIValue(m.CA),
		CustomCA:    stringToAPIValue(m.CustomCA),
		EABKID:      stringToAPIValue(m.EABKID),
		EABHMAC:     stringToAPIValue(m.EABHMAC),
	}
}

func (m *acmeclientAccountResourceModel) toAccountEditRequest() acmeclient.AccountEditRequest {
	return acmeclient.AccountEditRequest{
		Enabled:     boolToAPIString(m.Enabled),
		Name:        stringToAPIValue(m.Name),
		Description: stringToAPIValue(m.Description),
		Email:       stringToAPIValue(m.Email),
		CA:          stringToAPIValue(m.CA),
		CustomCA:    stringToAPIValue(m.CustomCA),
		EABKID:      stringToAPIValue(m.EABKID),
		EABHMAC:     stringToAPIValue(m.EABHMAC),
	}
}
