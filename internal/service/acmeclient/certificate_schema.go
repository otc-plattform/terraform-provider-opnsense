package acmeclient

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type acmeclientCertificateResourceModel struct {
	Id               types.String `tfsdk:"id"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	AltNames         types.Set    `tfsdk:"alt_names"`
	Account          types.String `tfsdk:"account_id"`
	ValidationMethod types.String `tfsdk:"validation_method_id"`
	AutoRenewal      types.Bool   `tfsdk:"auto_renewal"`
	RenewInterval    types.Int64  `tfsdk:"renew_interval"`
	KeyLength        types.String `tfsdk:"key_length"`
	OCSP             types.Bool   `tfsdk:"ocsp"`
	RestartActions   types.Set    `tfsdk:"restart_actions"`
	AliasMode        types.String `tfsdk:"alias_mode"`
	DomainAlias      types.String `tfsdk:"domain_alias"`
	ChallengeAlias   types.String `tfsdk:"challenge_alias"`
	CertificateRefId types.String `tfsdk:"certificate_ref_id"`
	LastUpdate       types.String `tfsdk:"last_update"`
	StatusCode       types.String `tfsdk:"status_code"`
	StatusLastUpdate types.String `tfsdk:"status_last_update"`
}

func acmeclientCertificateResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage ACME client certificates on OPNsense.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the certificate.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: "Whether the certificate is enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "Display name for the certificate.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "Optional description for the certificate.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"alt_names": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Subject Alternative Names for the certificate.",
				Optional:            true,
				Computed:            true,
			},
			"account_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the ACME account used for this certificate.",
				Required:            true,
			},
			"validation_method_id": schema.StringAttribute{
				MarkdownDescription: "UUID of the validation challenge used for this certificate.",
				Required:            true,
			},
			"auto_renewal": schema.BoolAttribute{
				MarkdownDescription: "Whether the certificate should auto-renew.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"renew_interval": schema.Int64Attribute{
				MarkdownDescription: "Renewal interval in days.",
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(60),
			},
			"key_length": schema.StringAttribute{
				MarkdownDescription: "Key length identifier (for example `key_2048`, `ec_p256`).",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("key_4096"),
				Validators: []validator.String{
					stringvalidator.OneOf("key_2048", "key_3072", "key_4096", "key_ec256", "key_ec384"),
				},
			},
			"ocsp": schema.BoolAttribute{
				MarkdownDescription: "Enable OCSP stapling.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
			},
			"restart_actions": schema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Services to restart after certificate issuance.",
				Optional:            true,
				Computed:            true,
			},
			"alias_mode": schema.StringAttribute{
				MarkdownDescription: "Alias mode for generated certificates.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("none"),
				Validators: []validator.String{
					stringvalidator.OneOf("none", "automatic", "domain", "challenge"),
				},
			},
			"domain_alias": schema.StringAttribute{
				MarkdownDescription: "Domain alias applied to the certificate.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"challenge_alias": schema.StringAttribute{
				MarkdownDescription: "Challenge alias applied to the certificate.",
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(""),
			},
			"certificate_ref_id": schema.StringAttribute{
				MarkdownDescription: "Reference ID of the generated certificate.",
				Computed:            true,
			},
			"last_update": schema.StringAttribute{
				MarkdownDescription: "Timestamp of the last certificate update.",
				Computed:            true,
			},
			"status_code": schema.StringAttribute{
				MarkdownDescription: "Last reported status code.",
				Computed:            true,
			},
			"status_last_update": schema.StringAttribute{
				MarkdownDescription: "Timestamp of the last status update.",
				Computed:            true,
			},
		},
	}
}

func acmeclientCertificateDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read ACME client certificate details from OPNsense.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the certificate.",
				Required:            true,
			},
			"enabled": dschema.BoolAttribute{
				MarkdownDescription: "Whether the certificate is enabled.",
				Computed:            true,
			},
			"name": dschema.StringAttribute{
				MarkdownDescription: "Display name for the certificate.",
				Computed:            true,
			},
			"description": dschema.StringAttribute{
				MarkdownDescription: "Description of the certificate.",
				Computed:            true,
			},
			"alt_names": dschema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Subject Alternative Names for the certificate.",
				Computed:            true,
			},
			"account_id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the ACME account used for this certificate.",
				Computed:            true,
			},
			"validation_method_id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the validation challenge used for this certificate.",
				Computed:            true,
			},
			"auto_renewal": dschema.BoolAttribute{
				MarkdownDescription: "Whether the certificate auto-renews.",
				Computed:            true,
			},
			"renew_interval": dschema.Int64Attribute{
				MarkdownDescription: "Renewal interval in days.",
				Computed:            true,
			},
			"key_length": dschema.StringAttribute{
				MarkdownDescription: "Key length identifier.",
				Computed:            true,
			},
			"ocsp": dschema.BoolAttribute{
				MarkdownDescription: "Whether OCSP stapling is enabled.",
				Computed:            true,
			},
			"restart_actions": dschema.SetAttribute{
				ElementType:         types.StringType,
				MarkdownDescription: "Services to restart after certificate issuance.",
				Computed:            true,
			},
			"alias_mode": dschema.StringAttribute{
				MarkdownDescription: "Alias mode for generated certificates.",
				Computed:            true,
			},
			"domain_alias": dschema.StringAttribute{
				MarkdownDescription: "Domain alias applied to the certificate.",
				Computed:            true,
			},
			"challenge_alias": dschema.StringAttribute{
				MarkdownDescription: "Challenge alias applied to the certificate.",
				Computed:            true,
			},
			"certificate_ref_id": dschema.StringAttribute{
				MarkdownDescription: "Reference ID of the generated certificate.",
				Computed:            true,
			},
			"last_update": dschema.StringAttribute{
				MarkdownDescription: "Timestamp of the last certificate update.",
				Computed:            true,
			},
			"status_code": dschema.StringAttribute{
				MarkdownDescription: "Last reported status code.",
				Computed:            true,
			},
			"status_last_update": dschema.StringAttribute{
				MarkdownDescription: "Timestamp of the last status update.",
				Computed:            true,
			},
		},
	}
}
