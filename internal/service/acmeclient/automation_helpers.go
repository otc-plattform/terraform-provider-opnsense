package acmeclient

import (
	"context"
	"errors"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/acmeclient"
	"github.com/browningluke/opnsense-go/pkg/errs"
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func automationResponseToModel(resp *acmeclient.ActionGetResponse) acmeclientAutomationResourceModel {
	action := resp.Action

	model := acmeclientAutomationResourceModel{
		Id:                    types.StringValue(action.ID),
		Enabled:               types.BoolValue(tools.StringToBool(action.Enabled)),
		Name:                  stringValueOrEmpty(action.Name),
		Description:           stringValueOrEmpty(action.Description),
		Type:                  types.StringValue(selectedOptionKey(action.Type)),
		SFTPHost:              stringValueOrEmpty(action.SFTPHost),
		SFTPPort:              stringValueOrEmpty(action.SFTPPort),
		SFTPHostKey:           stringValueOrEmpty(action.SFTPHostKey),
		SFTPUser:              stringValueOrEmpty(action.SFTPUser),
		SFTPIdentityType:      types.StringValue(selectedOptionKey(action.SFTPIdentityType)),
		SFTPRemotePath:        stringValueOrEmpty(action.SFTPRemotePath),
		SFTPChmod:             stringValueOrEmpty(action.SFTPChmod),
		SFTPChmodKey:          stringValueOrEmpty(action.SFTPChmodKey),
		SFTPChgrp:             stringValueOrEmpty(action.SFTPChgrp),
		SFTPFilenameCert:      stringValueOrEmpty(action.SFTPFilenameCert),
		SFTPFilenameKey:       stringValueOrEmpty(action.SFTPFilenameKey),
		SFTPFilenameCA:        stringValueOrEmpty(action.SFTPFilenameCA),
		SFTPFilenameFullchain: stringValueOrEmpty(action.SFTPFilenameFullchain),
		RemoteSSHHost:         stringValueOrEmpty(action.RemoteSSHHost),
		RemoteSSHPort:         stringValueOrEmpty(action.RemoteSSHPort),
		RemoteSSHKey:          stringValueOrEmpty(action.RemoteSSHHostKey),
		RemoteSSHUser:         stringValueOrEmpty(action.RemoteSSHUser),
		RemoteSSHIdentityType: types.StringValue(selectedOptionKey(action.RemoteSSHIdentityType)),
		RemoteSSHCommand:      stringValueOrEmpty(action.RemoteSSHCommand),
		ConfigdGenericCommand: types.StringValue(selectedOptionKey(action.ConfigdGenericCommand)),
		SynologyDSMHostname:   stringValueOrEmpty(action.SynologyDSMHostname),
		SynologyDSMPort:       stringValueOrEmpty(action.SynologyDSMPort),
		SynologyDSMScheme:     types.StringValue(selectedOptionKey(action.SynologyDSMScheme)),
		SynologyDSMUsername:   stringValueOrEmpty(action.SynologyDSMUsername),
		SynologyDSMPassword:   stringValueOrEmpty(action.SynologyDSMPassword),
		SynologyDSMDeviceID:   stringValueOrEmpty(action.SynologyDSMDeviceID),
		SynologyDSMDeviceName: stringValueOrEmpty(action.SynologyDSMDeviceName),
		SynologyDSMOTPCode:    stringValueOrEmpty(action.SynologyDSMOTPCode),
		SynologyDSMCreate:     types.BoolValue(tools.StringToBool(action.SynologyDSMCreate)),
		FritzboxURL:           stringValueOrEmpty(action.FritzBoxURL),
		FritzboxUsername:      stringValueOrEmpty(action.FritzBoxUsername),
		FritzboxPassword:      stringValueOrEmpty(action.FritzBoxPassword),
		PanosUsername:         stringValueOrEmpty(action.PANOSUsername),
		PanosPassword:         stringValueOrEmpty(action.PANOSPassword),
		PanosHost:             stringValueOrEmpty(action.PANOSHost),
		ProxmoxVEUser:         stringValueOrEmpty(action.ProxmoxVEUser),
		ProxmoxVEServer:       stringValueOrEmpty(action.ProxmoxVEServer),
		ProxmoxVEPort:         stringValueOrEmpty(action.ProxmoxVEPort),
		ProxmoxVENodeName:     stringValueOrEmpty(action.ProxmoxVENodeName),
		ProxmoxVERealm:        stringValueOrEmpty(action.ProxmoxVERealm),
		ProxmoxVETokenID:      stringValueOrEmpty(action.ProxmoxVETokenID),
		ProxmoxVETokenKey:     stringValueOrEmpty(action.ProxmoxVETokenKey),
		TruenasAPIKey:         stringValueOrEmpty(action.TrueNASAPIKey),
		TruenasHostname:       stringValueOrEmpty(action.TrueNASHostname),
		TruenasScheme:         types.StringValue(selectedOptionKey(action.TrueNASScheme)),
		UnifiKeystore:         stringValueOrEmpty(action.UnifiKeystore),
		VaultURL:              stringValueOrEmpty(action.VaultURL),
		VaultPrefix:           stringValueOrEmpty(action.VaultPrefix),
		VaultToken:            stringValueOrEmpty(action.VaultToken),
		VaultKVV2:             types.BoolValue(tools.StringToBool(action.VaultKVV2)),
	}

	return model
}

func fetchAutomationModel(ctx context.Context, controller *acmeclient.Controller, id string) (acmeclientAutomationResourceModel, error) {
	resp, err := controller.ACMEClientGetAutomation(ctx, id)
	if err == nil {
		model := automationResponseToModel(resp)
		if model.Id.IsNull() || model.Id.ValueString() == "" {
			model.Id = types.StringValue(id)
		}
		return model, nil
	}

	var notFound *errs.NotFoundError
	if errors.As(err, &notFound) {
		return acmeclientAutomationResourceModel{}, err
	}

	if !isAccountArrayDecodeError(err) {
		return acmeclientAutomationResourceModel{}, err
	}

	searchResp, searchErr := controller.ACMEClientSearchAutomation(ctx)
	if searchErr != nil {
		return acmeclientAutomationResourceModel{}, fmt.Errorf("fallback search failed: %w", searchErr)
	}

	for _, row := range searchResp.Rows {
		if row.UUID == id {
			return acmeclientAutomationResourceModel{
				Id:          types.StringValue(row.UUID),
				Enabled:     types.BoolValue(tools.StringToBool(row.Enabled)),
				Name:        stringValueOrNull(row.Name),
				Description: stringValueOrNull(row.Description),
				Type:        stringValueOrNull(row.Type),
			}, nil
		}
	}

	return acmeclientAutomationResourceModel{}, errs.NewNotFoundError()
}

func (m *acmeclientAutomationResourceModel) toAction() acmeclient.Action {
	action := acmeclient.Action{
		Enabled:               boolToAPIString(m.Enabled),
		Name:                  stringToAPIValue(m.Name),
		Description:           stringToAPIValue(m.Description),
		Type:                  stringToAPIValue(m.Type),
		SFTPHost:              stringToAPIValue(m.SFTPHost),
		SFTPPort:              stringToAPIValue(m.SFTPPort),
		SFTPHostKey:           stringToAPIValue(m.SFTPHostKey),
		SFTPUser:              stringToAPIValue(m.SFTPUser),
		SFTPIdentityType:      stringToAPIValue(m.SFTPIdentityType),
		SFTPRemotePath:        stringToAPIValue(m.SFTPRemotePath),
		SFTPChmod:             stringToAPIValue(m.SFTPChmod),
		SFTPChmodKey:          stringToAPIValue(m.SFTPChmodKey),
		SFTPChgrp:             stringToAPIValue(m.SFTPChgrp),
		SFTPFilenameCert:      stringToAPIValue(m.SFTPFilenameCert),
		SFTPFilenameKey:       stringToAPIValue(m.SFTPFilenameKey),
		SFTPFilenameCA:        stringToAPIValue(m.SFTPFilenameCA),
		SFTPFilenameFullchain: stringToAPIValue(m.SFTPFilenameFullchain),
		RemoteSSHHost:         stringToAPIValue(m.RemoteSSHHost),
		RemoteSSHPort:         stringToAPIValue(m.RemoteSSHPort),
		RemoteSSHKey:          stringToAPIValue(m.RemoteSSHKey),
		RemoteSSHUser:         stringToAPIValue(m.RemoteSSHUser),
		RemoteSSHIdentityType: stringToAPIValue(m.RemoteSSHIdentityType),
		RemoteSSHCommand:      stringToAPIValue(m.RemoteSSHCommand),
		ConfigdGenericCommand: stringToAPIValue(m.ConfigdGenericCommand),
		SynologyDSMHostname:   stringToAPIValue(m.SynologyDSMHostname),
		SynologyDSMPort:       stringToAPIValue(m.SynologyDSMPort),
		SynologyDSMScheme:     stringToAPIValue(m.SynologyDSMScheme),
		SynologyDSMUsername:   stringToAPIValue(m.SynologyDSMUsername),
		SynologyDSMPassword:   stringToAPIValue(m.SynologyDSMPassword),
		SynologyDSMDeviceID:   stringToAPIValue(m.SynologyDSMDeviceID),
		SynologyDSMDeviceName: stringToAPIValue(m.SynologyDSMDeviceName),
		SynologyDSMOTPCode:    stringToAPIValue(m.SynologyDSMOTPCode),
		SynologyDSMCreate:     boolToAPIString(m.SynologyDSMCreate),
		FritzBoxURL:           stringToAPIValue(m.FritzboxURL),
		FritzBoxUsername:      stringToAPIValue(m.FritzboxUsername),
		FritzBoxPassword:      stringToAPIValue(m.FritzboxPassword),
		PANOSUsername:         stringToAPIValue(m.PanosUsername),
		PANOSPassword:         stringToAPIValue(m.PanosPassword),
		PANOSHost:             stringToAPIValue(m.PanosHost),
		ProxmoxVEUser:         stringToAPIValue(m.ProxmoxVEUser),
		ProxmoxVEServer:       stringToAPIValue(m.ProxmoxVEServer),
		ProxmoxVEPort:         stringToAPIValue(m.ProxmoxVEPort),
		ProxmoxVENodeName:     stringToAPIValue(m.ProxmoxVENodeName),
		ProxmoxVERealm:        stringToAPIValue(m.ProxmoxVERealm),
		ProxmoxVETokenID:      stringToAPIValue(m.ProxmoxVETokenID),
		ProxmoxVETokenKey:     stringToAPIValue(m.ProxmoxVETokenKey),
		TrueNASAPIKey:         stringToAPIValue(m.TruenasAPIKey),
		TrueNASHostname:       stringToAPIValue(m.TruenasHostname),
		TrueNASScheme:         stringToAPIValue(m.TruenasScheme),
		UnifiKeystore:         stringToAPIValue(m.UnifiKeystore),
		VaultURL:              stringToAPIValue(m.VaultURL),
		VaultPrefix:           stringToAPIValue(m.VaultPrefix),
		VaultToken:            stringToAPIValue(m.VaultToken),
		VaultKVV2:             boolToAPIString(m.VaultKVV2),
	}

	return action
}

func mergeAutomationSensitiveFields(state *acmeclientAutomationResourceModel, prior *acmeclientAutomationResourceModel) {
	if state == nil || prior == nil {
		return
	}

	copyStringIfEmpty(&state.SFTPHostKey, prior.SFTPHostKey)
	copyStringIfEmpty(&state.RemoteSSHKey, prior.RemoteSSHKey)
	copyStringIfEmpty(&state.SynologyDSMPassword, prior.SynologyDSMPassword)
	copyStringIfEmpty(&state.FritzboxPassword, prior.FritzboxPassword)
	copyStringIfEmpty(&state.PanosPassword, prior.PanosPassword)
	copyStringIfEmpty(&state.ProxmoxVETokenKey, prior.ProxmoxVETokenKey)
	copyStringIfEmpty(&state.TruenasAPIKey, prior.TruenasAPIKey)
	copyStringIfEmpty(&state.VaultToken, prior.VaultToken)
}

func copyStringIfEmpty(target *types.String, source types.String) {
	if target == nil {
		return
	}
	if !target.IsNull() && !target.IsUnknown() && target.ValueString() != "" {
		return
	}
	if source.IsNull() || source.IsUnknown() {
		return
	}
	if source.ValueString() == "" {
		return
	}
	*target = source
}

func stringValueOrEmpty(s string) types.String {
	return types.StringValue(s)
}
