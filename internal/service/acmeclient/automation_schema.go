package acmeclient

import (
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type acmeclientAutomationResourceModel struct {
	Id                    types.String `tfsdk:"id"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	Name                  types.String `tfsdk:"name"`
	Description           types.String `tfsdk:"description"`
	Type                  types.String `tfsdk:"type"`
	SFTPHost              types.String `tfsdk:"sftp_host"`
	SFTPPort              types.String `tfsdk:"sftp_port"`
	SFTPHostKey           types.String `tfsdk:"sftp_host_key"`
	SFTPUser              types.String `tfsdk:"sftp_user"`
	SFTPIdentityType      types.String `tfsdk:"sftp_identity_type"`
	SFTPRemotePath        types.String `tfsdk:"sftp_remote_path"`
	SFTPChmod             types.String `tfsdk:"sftp_chmod"`
	SFTPChmodKey          types.String `tfsdk:"sftp_chmod_key"`
	SFTPChgrp             types.String `tfsdk:"sftp_chgrp"`
	SFTPFilenameCert      types.String `tfsdk:"sftp_filename_cert"`
	SFTPFilenameKey       types.String `tfsdk:"sftp_filename_key"`
	SFTPFilenameCA        types.String `tfsdk:"sftp_filename_ca"`
	SFTPFilenameFullchain types.String `tfsdk:"sftp_filename_fullchain"`
	RemoteSSHHost         types.String `tfsdk:"remote_ssh_host"`
	RemoteSSHPort         types.String `tfsdk:"remote_ssh_port"`
	RemoteSSHKey          types.String `tfsdk:"remote_ssh_key"`
	RemoteSSHUser         types.String `tfsdk:"remote_ssh_user"`
	RemoteSSHIdentityType types.String `tfsdk:"remote_ssh_identity_type"`
	RemoteSSHCommand      types.String `tfsdk:"remote_ssh_command"`
	ConfigdGenericCommand types.String `tfsdk:"configd_generic_command"`
	SynologyDSMHostname   types.String `tfsdk:"synology_dsm_hostname"`
	SynologyDSMPort       types.String `tfsdk:"synology_dsm_port"`
	SynologyDSMScheme     types.String `tfsdk:"synology_dsm_scheme"`
	SynologyDSMUsername   types.String `tfsdk:"synology_dsm_username"`
	SynologyDSMPassword   types.String `tfsdk:"synology_dsm_password"`
	SynologyDSMDeviceID   types.String `tfsdk:"synology_dsm_device_id"`
	SynologyDSMDeviceName types.String `tfsdk:"synology_dsm_device_name"`
	SynologyDSMOTPCode    types.String `tfsdk:"synology_dsm_otp_code"`
	SynologyDSMCreate     types.Bool   `tfsdk:"synology_dsm_create"`
	FritzboxURL           types.String `tfsdk:"fritzbox_url"`
	FritzboxUsername      types.String `tfsdk:"fritzbox_username"`
	FritzboxPassword      types.String `tfsdk:"fritzbox_password"`
	PanosUsername         types.String `tfsdk:"panos_username"`
	PanosPassword         types.String `tfsdk:"panos_password"`
	PanosHost             types.String `tfsdk:"panos_host"`
	ProxmoxVEUser         types.String `tfsdk:"proxmoxve_user"`
	ProxmoxVEServer       types.String `tfsdk:"proxmoxve_server"`
	ProxmoxVEPort         types.String `tfsdk:"proxmoxve_port"`
	ProxmoxVENodeName     types.String `tfsdk:"proxmoxve_node_name"`
	ProxmoxVERealm        types.String `tfsdk:"proxmoxve_realm"`
	ProxmoxVETokenID      types.String `tfsdk:"proxmoxve_token_id"`
	ProxmoxVETokenKey     types.String `tfsdk:"proxmoxve_token_key"`
	TruenasAPIKey         types.String `tfsdk:"truenas_api_key"`
	TruenasHostname       types.String `tfsdk:"truenas_hostname"`
	TruenasScheme         types.String `tfsdk:"truenas_scheme"`
	UnifiKeystore         types.String `tfsdk:"unifi_keystore"`
	VaultURL              types.String `tfsdk:"vault_url"`
	VaultPrefix           types.String `tfsdk:"vault_prefix"`
	VaultToken            types.String `tfsdk:"vault_token"`
	VaultKVV2             types.Bool   `tfsdk:"vault_kvv2"`
}

func acmeclientAutomationResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage ACME client automation actions on OPNsense.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "UUID of the automation action.",
			},
			"enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Whether the automation is enabled.",
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Display name of the automation action.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
				MarkdownDescription: "Optional description of the automation action.",
			},
			"type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Automation action type (e.g. `configd_restart_gui`, `configd_upload_sftp`, `acme_proxmoxve`).",
			},
			"sftp_host": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_port": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("22"),
			},
			"sftp_host_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_user": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_identity_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_remote_path": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_chmod": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_chmod_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_chgrp": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_filename_cert": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_filename_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_filename_ca": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"sftp_filename_fullchain": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"remote_ssh_host": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"remote_ssh_port": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("22"),
			},
			"remote_ssh_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"remote_ssh_user": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"remote_ssh_identity_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"remote_ssh_command": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"configd_generic_command": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"synology_dsm_hostname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"synology_dsm_port": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("5000"),
			},
			"synology_dsm_scheme": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("http"),
			},
			"synology_dsm_username": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"synology_dsm_password": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"synology_dsm_device_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"synology_dsm_device_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"synology_dsm_otp_code": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"synology_dsm_create": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
			"fritzbox_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"fritzbox_username": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"fritzbox_password": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"panos_username": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"panos_password": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"panos_host": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"proxmoxve_user": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("root"),
			},
			"proxmoxve_server": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"proxmoxve_port": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("8006"),
			},
			"proxmoxve_node_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"proxmoxve_realm": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("pam"),
			},
			"proxmoxve_token_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("acme"),
			},
			"proxmoxve_token_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"truenas_api_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"truenas_hostname": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("localhost"),
			},
			"truenas_scheme": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("http"),
			},
			"unifi_keystore": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("/usr/local/share/java/unifi/data/keystore"),
			},
			"vault_url": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"vault_prefix": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("acme"),
			},
			"vault_token": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString(""),
			},
			"vault_kvv2": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Default:  booldefault.StaticBool(true),
			},
		},
	}
}

func acmeclientAutomationDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read ACME client automation details from OPNsense.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the automation action.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether the automation is enabled.",
				Computed:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Display name of the automation action.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description of the automation action.",
				Computed:            true,
			},
			"type": dschema.StringAttribute{
				MarkdownDescription: "Automation action type.",
				Computed:            true,
			},
			"sftp_host": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Target host for SFTP uploads.",
			},
			"sftp_port": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SFTP port.",
			},
			"sftp_host_key": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SSH host key fingerprint for SFTP uploads.",
			},
			"sftp_user": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SFTP username.",
			},
			"sftp_identity_type": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SFTP identity type.",
			},
			"sftp_remote_path": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Remote path used during SFTP uploads.",
			},
			"sftp_chmod": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "File permission applied to SFTP uploads.",
			},
			"sftp_chmod_key": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Key file permission applied to SFTP uploads.",
			},
			"sftp_chgrp": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Group applied to SFTP uploads.",
			},
			"sftp_filename_cert": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Filename for the certificate uploaded via SFTP.",
			},
			"sftp_filename_key": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Filename for the private key uploaded via SFTP.",
			},
			"sftp_filename_ca": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Filename for the CA bundle uploaded via SFTP.",
			},
			"sftp_filename_fullchain": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Filename for the full chain uploaded via SFTP.",
			},
			"remote_ssh_host": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Remote SSH host for command execution.",
			},
			"remote_ssh_port": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Remote SSH port.",
			},
			"remote_ssh_key": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "SSH host key fingerprint for remote execution.",
			},
			"remote_ssh_user": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Remote SSH user.",
			},
			"remote_ssh_identity_type": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Identity type for remote SSH authentication.",
			},
			"remote_ssh_command": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Command executed on the remote host via SSH.",
			},
			"configd_generic_command": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Configd command executed as part of the automation.",
			},
			"synology_dsm_hostname": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM hostname.",
			},
			"synology_dsm_port": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM port.",
			},
			"synology_dsm_scheme": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM scheme (http or https).",
			},
			"synology_dsm_username": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM user.",
			},
			"synology_dsm_password": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM password.",
			},
			"synology_dsm_device_id": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM device identifier.",
			},
			"synology_dsm_device_name": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM device name.",
			},
			"synology_dsm_otp_code": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Synology DSM OTP code.",
			},
			"synology_dsm_create": dschema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether the Synology DSM host should be created if missing.",
			},
			"fritzbox_url": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "FRITZ!Box router URL.",
			},
			"fritzbox_username": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "FRITZ!Box username.",
			},
			"fritzbox_password": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "FRITZ!Box password.",
			},
			"panos_username": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Palo Alto Networks username.",
			},
			"panos_password": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Palo Alto Networks password.",
			},
			"panos_host": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Palo Alto Networks host.",
			},
			"proxmoxve_user": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Proxmox VE user.",
			},
			"proxmoxve_server": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Proxmox VE server.",
			},
			"proxmoxve_port": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Proxmox VE API port.",
			},
			"proxmoxve_node_name": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Proxmox VE node name.",
			},
			"proxmoxve_realm": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Proxmox VE realm.",
			},
			"proxmoxve_token_id": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Proxmox VE token ID.",
			},
			"proxmoxve_token_key": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Proxmox VE token key.",
			},
			"truenas_api_key": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "TrueNAS API key.",
			},
			"truenas_hostname": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "TrueNAS hostname.",
			},
			"truenas_scheme": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "TrueNAS scheme (`http` or `https`).",
			},
			"unifi_keystore": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Path to the Unifi keystore.",
			},
			"vault_url": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Hashicorp Vault API URL.",
			},
			"vault_prefix": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Vault KV secret prefix.",
			},
			"vault_token": dschema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "Vault token used for authentication.",
			},
			"vault_kvv2": dschema.BoolAttribute{
				Computed:            true,
				MarkdownDescription: "Whether Vault KV v2 is used.",
			},
		},
	}
}
