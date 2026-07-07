package haproxy

import (
	"github.com/browningluke/terraform-provider-opnsense/internal/tools"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	dschema "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type aclResourceModel struct {
	Id types.String `tfsdk:"id"`
	InternalId types.String `tfsdk:"internal_id"`
	Name types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Expression types.String `tfsdk:"expression"`
	Negate types.Bool `tfsdk:"negate"`
	CaseSensitive types.Bool `tfsdk:"case_sensitive"`
	HdrBeg types.String `tfsdk:"hdr_beg"`
	HdrEnd types.String `tfsdk:"hdr_end"`
	Hdr types.String `tfsdk:"hdr"`
	HdrReg types.String `tfsdk:"hdr_reg"`
	HdrSub types.String `tfsdk:"hdr_sub"`
	PathBeg types.String `tfsdk:"path_beg"`
	PathEnd types.String `tfsdk:"path_end"`
	Path types.String `tfsdk:"path"`
	PathReg types.String `tfsdk:"path_reg"`
	PathDir types.String `tfsdk:"path_dir"`
	PathSub types.String `tfsdk:"path_sub"`
	CustHdrBegName types.String `tfsdk:"cust_hdr_beg_name"`
	CustHdrBeg types.String `tfsdk:"cust_hdr_beg"`
	CustHdrEndName types.String `tfsdk:"cust_hdr_end_name"`
	CustHdrEnd types.String `tfsdk:"cust_hdr_end"`
	CustHdrName types.String `tfsdk:"cust_hdr_name"`
	CustHdr types.String `tfsdk:"cust_hdr"`
	CustHdrRegName types.String `tfsdk:"cust_hdr_reg_name"`
	CustHdrReg types.String `tfsdk:"cust_hdr_reg"`
	CustHdrSubName types.String `tfsdk:"cust_hdr_sub_name"`
	CustHdrSub types.String `tfsdk:"cust_hdr_sub"`
	UrlParam types.String `tfsdk:"url_param"`
	UrlParamValue types.String `tfsdk:"url_param_value"`
	Var_ types.String `tfsdk:"var"`
	VarValue types.String `tfsdk:"var_value"`
	VarComparison types.String `tfsdk:"var_comparison"`
	SslCVerifyCode types.String `tfsdk:"ssl_c_verify_code"`
	SslCCaCommonname types.String `tfsdk:"ssl_c_ca_commonname"`
	SslHelloType types.String `tfsdk:"ssl_hello_type"`
	Src types.String `tfsdk:"src"`
	SrcBytesInRateComparison types.String `tfsdk:"src_bytes_in_rate_comparison"`
	SrcBytesInRate types.String `tfsdk:"src_bytes_in_rate"`
	SrcBytesOutRateComparison types.String `tfsdk:"src_bytes_out_rate_comparison"`
	SrcBytesOutRate types.String `tfsdk:"src_bytes_out_rate"`
	SrcConnCntComparison types.String `tfsdk:"src_conn_cnt_comparison"`
	SrcConnCnt types.String `tfsdk:"src_conn_cnt"`
	SrcConnCurComparison types.String `tfsdk:"src_conn_cur_comparison"`
	SrcConnCur types.String `tfsdk:"src_conn_cur"`
	SrcConnRateComparison types.String `tfsdk:"src_conn_rate_comparison"`
	SrcConnRate types.String `tfsdk:"src_conn_rate"`
	SrcHttpErrCntComparison types.String `tfsdk:"src_http_err_cnt_comparison"`
	SrcHttpErrCnt types.String `tfsdk:"src_http_err_cnt"`
	SrcHttpErrRateComparison types.String `tfsdk:"src_http_err_rate_comparison"`
	SrcHttpErrRate types.String `tfsdk:"src_http_err_rate"`
	SrcHttpReqCntComparison types.String `tfsdk:"src_http_req_cnt_comparison"`
	SrcHttpReqCnt types.String `tfsdk:"src_http_req_cnt"`
	SrcHttpReqRateComparison types.String `tfsdk:"src_http_req_rate_comparison"`
	SrcHttpReqRate types.String `tfsdk:"src_http_req_rate"`
	SrcKbytesInComparison types.String `tfsdk:"src_kbytes_in_comparison"`
	SrcKbytesIn types.String `tfsdk:"src_kbytes_in"`
	SrcKbytesOutComparison types.String `tfsdk:"src_kbytes_out_comparison"`
	SrcKbytesOut types.String `tfsdk:"src_kbytes_out"`
	SrcPortComparison types.String `tfsdk:"src_port_comparison"`
	SrcPort types.String `tfsdk:"src_port"`
	SrcSessCntComparison types.String `tfsdk:"src_sess_cnt_comparison"`
	SrcSessCnt types.String `tfsdk:"src_sess_cnt"`
	SrcSessRateComparison types.String `tfsdk:"src_sess_rate_comparison"`
	SrcSessRate types.String `tfsdk:"src_sess_rate"`
	Nbsrv types.String `tfsdk:"nbsrv"`
	NbsrvBackend types.String `tfsdk:"nbsrv_backend"`
	SslFcSni types.String `tfsdk:"ssl_fc_sni"`
	SslSni types.String `tfsdk:"ssl_sni"`
	SslSniSub types.String `tfsdk:"ssl_sni_sub"`
	SslSniBeg types.String `tfsdk:"ssl_sni_beg"`
	SslSniEnd types.String `tfsdk:"ssl_sni_end"`
	SslSniReg types.String `tfsdk:"ssl_sni_reg"`
	CustomAcl types.String `tfsdk:"custom_acl"`
	Value types.String `tfsdk:"value"`
	Urlparam types.String `tfsdk:"urlparam"`
	QueryBackend types.String `tfsdk:"query_backend"`
	AllowedUsers types.Set `tfsdk:"allowed_users"`
	AllowedGroups types.Set `tfsdk:"allowed_groups"`
	HttpMethod types.Set `tfsdk:"http_method"`
	ScBytesInRateComparison types.String `tfsdk:"sc_bytes_in_rate_comparison"`
	ScBytesInRate types.String `tfsdk:"sc_bytes_in_rate"`
	ScBytesOutRateComparison types.String `tfsdk:"sc_bytes_out_rate_comparison"`
	ScBytesOutRate types.String `tfsdk:"sc_bytes_out_rate"`
	ScClrGpcComparison types.String `tfsdk:"sc_clr_gpc_comparison"`
	ScClrGpc types.String `tfsdk:"sc_clr_gpc"`
	ScConnCntComparison types.String `tfsdk:"sc_conn_cnt_comparison"`
	ScConnCnt types.String `tfsdk:"sc_conn_cnt"`
	ScConnCurComparison types.String `tfsdk:"sc_conn_cur_comparison"`
	ScConnCur types.String `tfsdk:"sc_conn_cur"`
	ScConnRateComparison types.String `tfsdk:"sc_conn_rate_comparison"`
	ScConnRate types.String `tfsdk:"sc_conn_rate"`
	ScGetGpcComparison types.String `tfsdk:"sc_get_gpc_comparison"`
	ScGetGpc types.String `tfsdk:"sc_get_gpc"`
	ScGlitchCntComparison types.String `tfsdk:"sc_glitch_cnt_comparison"`
	ScGlitchCnt types.String `tfsdk:"sc_glitch_cnt"`
	ScGlitchRateComparison types.String `tfsdk:"sc_glitch_rate_comparison"`
	ScGlitchRate types.String `tfsdk:"sc_glitch_rate"`
	ScGpcRateComparison types.String `tfsdk:"sc_gpc_rate_comparison"`
	ScGpcRate types.String `tfsdk:"sc_gpc_rate"`
	ScHttpErrCntComparison types.String `tfsdk:"sc_http_err_cnt_comparison"`
	ScHttpErrCnt types.String `tfsdk:"sc_http_err_cnt"`
	ScHttpErrRateComparison types.String `tfsdk:"sc_http_err_rate_comparison"`
	ScHttpErrRate types.String `tfsdk:"sc_http_err_rate"`
	ScHttpFailCntComparison types.String `tfsdk:"sc_http_fail_cnt_comparison"`
	ScHttpFailCnt types.String `tfsdk:"sc_http_fail_cnt"`
	ScHttpFailRateComparison types.String `tfsdk:"sc_http_fail_rate_comparison"`
	ScHttpFailRate types.String `tfsdk:"sc_http_fail_rate"`
	ScHttpReqCntComparison types.String `tfsdk:"sc_http_req_cnt_comparison"`
	ScHttpReqCnt types.String `tfsdk:"sc_http_req_cnt"`
	ScHttpReqRateComparison types.String `tfsdk:"sc_http_req_rate_comparison"`
	ScHttpReqRate types.String `tfsdk:"sc_http_req_rate"`
	ScIncGpcComparison types.String `tfsdk:"sc_inc_gpc_comparison"`
	ScIncGpc types.String `tfsdk:"sc_inc_gpc"`
	ScSessCntComparison types.String `tfsdk:"sc_sess_cnt_comparison"`
	ScSessCnt types.String `tfsdk:"sc_sess_cnt"`
	ScSessRateComparison types.String `tfsdk:"sc_sess_rate_comparison"`
	ScSessRate types.String `tfsdk:"sc_sess_rate"`
	SrcGetGpcComparison types.String `tfsdk:"src_get_gpc_comparison"`
	SrcGetGpc types.String `tfsdk:"src_get_gpc"`
	SrcGetGptComparison types.String `tfsdk:"src_get_gpt_comparison"`
	SrcGetGpt types.String `tfsdk:"src_get_gpt"`
	SrcGlitchCntComparison types.String `tfsdk:"src_glitch_cnt_comparison"`
	SrcGlitchCnt types.String `tfsdk:"src_glitch_cnt"`
	SrcGlitchRateComparison types.String `tfsdk:"src_glitch_rate_comparison"`
	SrcGlitchRate types.String `tfsdk:"src_glitch_rate"`
	SrcGpcRateComparison types.String `tfsdk:"src_gpc_rate_comparison"`
	SrcGpcRate types.String `tfsdk:"src_gpc_rate"`
	SrcHttpFailCntComparison types.String `tfsdk:"src_http_fail_cnt_comparison"`
	SrcHttpFailCnt types.String `tfsdk:"src_http_fail_cnt"`
	SrcHttpFailRateComparison types.String `tfsdk:"src_http_fail_rate_comparison"`
	SrcHttpFailRate types.String `tfsdk:"src_http_fail_rate"`
	SrcIncGpcComparison types.String `tfsdk:"src_inc_gpc_comparison"`
	SrcIncGpc types.String `tfsdk:"src_inc_gpc"`
	ScClrGpc0Comparison types.String `tfsdk:"sc_clr_gpc0_comparison"`
	ScClrGpc0 types.String `tfsdk:"sc_clr_gpc0"`
	ScClrGpc1Comparison types.String `tfsdk:"sc_clr_gpc1_comparison"`
	ScClrGpc1 types.String `tfsdk:"sc_clr_gpc1"`
	Sc0ClrGpc0Comparison types.String `tfsdk:"sc0_clr_gpc0_comparison"`
	Sc0ClrGpc0 types.String `tfsdk:"sc0_clr_gpc0"`
	Sc0ClrGpc1Comparison types.String `tfsdk:"sc0_clr_gpc1_comparison"`
	Sc0ClrGpc1 types.String `tfsdk:"sc0_clr_gpc1"`
	Sc1ClrGpcComparison types.String `tfsdk:"sc1_clr_gpc_comparison"`
	Sc1ClrGpc types.String `tfsdk:"sc1_clr_gpc"`
	Sc1ClrGpc0Comparison types.String `tfsdk:"sc1_clr_gpc0_comparison"`
	Sc1ClrGpc0 types.String `tfsdk:"sc1_clr_gpc0"`
	Sc1ClrGpc1Comparison types.String `tfsdk:"sc1_clr_gpc1_comparison"`
	Sc1ClrGpc1 types.String `tfsdk:"sc1_clr_gpc1"`
	Sc2ClrGpcComparison types.String `tfsdk:"sc2_clr_gpc_comparison"`
	Sc2ClrGpc types.String `tfsdk:"sc2_clr_gpc"`
	Sc2ClrGpc0Comparison types.String `tfsdk:"sc2_clr_gpc0_comparison"`
	Sc2ClrGpc0 types.String `tfsdk:"sc2_clr_gpc0"`
	Sc2ClrGpc1Comparison types.String `tfsdk:"sc2_clr_gpc1_comparison"`
	Sc2ClrGpc1 types.String `tfsdk:"sc2_clr_gpc1"`
	ScGetGpc0Comparison types.String `tfsdk:"sc_get_gpc0_comparison"`
	ScGetGpc0 types.String `tfsdk:"sc_get_gpc0"`
	ScGetGpc1Comparison types.String `tfsdk:"sc_get_gpc1_comparison"`
	ScGetGpc1 types.String `tfsdk:"sc_get_gpc1"`
	Sc0GetGpc0Comparison types.String `tfsdk:"sc0_get_gpc0_comparison"`
	Sc0GetGpc0 types.String `tfsdk:"sc0_get_gpc0"`
	Sc0GetGpc1Comparison types.String `tfsdk:"sc0_get_gpc1_comparison"`
	Sc0GetGpc1 types.String `tfsdk:"sc0_get_gpc1"`
	Sc1GetGpc0Comparison types.String `tfsdk:"sc1_get_gpc0_comparison"`
	Sc1GetGpc0 types.String `tfsdk:"sc1_get_gpc0"`
	Sc1GetGpc1Comparison types.String `tfsdk:"sc1_get_gpc1_comparison"`
	Sc1GetGpc1 types.String `tfsdk:"sc1_get_gpc1"`
	Sc2GetGpc0Comparison types.String `tfsdk:"sc2_get_gpc0_comparison"`
	Sc2GetGpc0 types.String `tfsdk:"sc2_get_gpc0"`
	Sc2GetGpc1Comparison types.String `tfsdk:"sc2_get_gpc1_comparison"`
	Sc2GetGpc1 types.String `tfsdk:"sc2_get_gpc1"`
	ScGetGptComparison types.String `tfsdk:"sc_get_gpt_comparison"`
	ScGetGpt types.String `tfsdk:"sc_get_gpt"`
	ScGetGpt0Comparison types.String `tfsdk:"sc_get_gpt0_comparison"`
	ScGetGpt0 types.String `tfsdk:"sc_get_gpt0"`
	Sc0GetGpt0Comparison types.String `tfsdk:"sc0_get_gpt0_comparison"`
	Sc0GetGpt0 types.String `tfsdk:"sc0_get_gpt0"`
	Sc1GetGpt0Comparison types.String `tfsdk:"sc1_get_gpt0_comparison"`
	Sc1GetGpt0 types.String `tfsdk:"sc1_get_gpt0"`
	Sc2GetGpt0Comparison types.String `tfsdk:"sc2_get_gpt0_comparison"`
	Sc2GetGpt0 types.String `tfsdk:"sc2_get_gpt0"`
	ScGpc0RateComparison types.String `tfsdk:"sc_gpc0_rate_comparison"`
	ScGpc0Rate types.String `tfsdk:"sc_gpc0_rate"`
	ScGpc1RateComparison types.String `tfsdk:"sc_gpc1_rate_comparison"`
	ScGpc1Rate types.String `tfsdk:"sc_gpc1_rate"`
	Sc0Gpc0RateComparison types.String `tfsdk:"sc0_gpc0_rate_comparison"`
	Sc0Gpc0Rate types.String `tfsdk:"sc0_gpc0_rate"`
	Sc0Gpc1RateComparison types.String `tfsdk:"sc0_gpc1_rate_comparison"`
	Sc0Gpc1Rate types.String `tfsdk:"sc0_gpc1_rate"`
	Sc1Gpc0RateComparison types.String `tfsdk:"sc1_gpc0_rate_comparison"`
	Sc1Gpc0Rate types.String `tfsdk:"sc1_gpc0_rate"`
	Sc1Gpc1RateComparison types.String `tfsdk:"sc1_gpc1_rate_comparison"`
	Sc1Gpc1Rate types.String `tfsdk:"sc1_gpc1_rate"`
	Sc2Gpc0RateComparison types.String `tfsdk:"sc2_gpc0_rate_comparison"`
	Sc2Gpc0Rate types.String `tfsdk:"sc2_gpc0_rate"`
	Sc2Gpc1RateComparison types.String `tfsdk:"sc2_gpc1_rate_comparison"`
	Sc2Gpc1Rate types.String `tfsdk:"sc2_gpc1_rate"`
	ScIncGpc0Comparison types.String `tfsdk:"sc_inc_gpc0_comparison"`
	ScIncGpc0 types.String `tfsdk:"sc_inc_gpc0"`
	ScIncGpc1Comparison types.String `tfsdk:"sc_inc_gpc1_comparison"`
	ScIncGpc1 types.String `tfsdk:"sc_inc_gpc1"`
	Sc0IncGpc0Comparison types.String `tfsdk:"sc0_inc_gpc0_comparison"`
	Sc0IncGpc0 types.String `tfsdk:"sc0_inc_gpc0"`
	Sc0IncGpc1Comparison types.String `tfsdk:"sc0_inc_gpc1_comparison"`
	Sc0IncGpc1 types.String `tfsdk:"sc0_inc_gpc1"`
	Sc1IncGpc0Comparison types.String `tfsdk:"sc1_inc_gpc0_comparison"`
	Sc1IncGpc0 types.String `tfsdk:"sc1_inc_gpc0"`
	Sc1IncGpc1Comparison types.String `tfsdk:"sc1_inc_gpc1_comparison"`
	Sc1IncGpc1 types.String `tfsdk:"sc1_inc_gpc1"`
	Sc2IncGpc0Comparison types.String `tfsdk:"sc2_inc_gpc0_comparison"`
	Sc2IncGpc0 types.String `tfsdk:"sc2_inc_gpc0"`
	Sc2IncGpc1Comparison types.String `tfsdk:"sc2_inc_gpc1_comparison"`
	Sc2IncGpc1 types.String `tfsdk:"sc2_inc_gpc1"`
	SrcClrGpc0Comparison types.String `tfsdk:"src_clr_gpc0_comparison"`
	SrcClrGpc0 types.String `tfsdk:"src_clr_gpc0"`
	SrcClrGpc1Comparison types.String `tfsdk:"src_clr_gpc1_comparison"`
	SrcClrGpc1 types.String `tfsdk:"src_clr_gpc1"`
	SrcGetGpc0Comparison types.String `tfsdk:"src_get_gpc0_comparison"`
	SrcGetGpc0 types.String `tfsdk:"src_get_gpc0"`
	SrcGetGpc1Comparison types.String `tfsdk:"src_get_gpc1_comparison"`
	SrcGetGpc1 types.String `tfsdk:"src_get_gpc1"`
	SrcGpc0RateComparison types.String `tfsdk:"src_gpc0_rate_comparison"`
	SrcGpc0Rate types.String `tfsdk:"src_gpc0_rate"`
	SrcGpc1RateComparison types.String `tfsdk:"src_gpc1_rate_comparison"`
	SrcGpc1Rate types.String `tfsdk:"src_gpc1_rate"`
	SrcIncGpc0Comparison types.String `tfsdk:"src_inc_gpc0_comparison"`
	SrcIncGpc0 types.String `tfsdk:"src_inc_gpc0"`
	SrcIncGpc1Comparison types.String `tfsdk:"src_inc_gpc1_comparison"`
	SrcIncGpc1 types.String `tfsdk:"src_inc_gpc1"`
	GpcNumber types.String `tfsdk:"gpc_number"`
	GptNumber types.String `tfsdk:"gpt_number"`
	ScNumber types.String `tfsdk:"sc_number"`
	TableName types.String `tfsdk:"table_name"`
	Mapfile types.String `tfsdk:"mapfile"`
	Converter types.String `tfsdk:"converter"`
}

func aclResourceSchema() schema.Schema {
	return schema.Schema{
		MarkdownDescription: "Manage an OPNsense HAProxy condition/ACL.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy acl.",
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"internal_id": schema.StringAttribute{
				MarkdownDescription: "Internal OPNsense id assigned to this haproxy acl.",
				Computed: true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "ACL name value.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "ACL description value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"expression": schema.StringAttribute{
				MarkdownDescription: "ACL expression option. One of: cust_hdr_beg, cust_hdr_end, cust_hdr, cust_hdr_reg, cust_hdr_sub, hdr_beg, hdr_end, hdr, hdr_reg, hdr_sub, http_auth, http_method, nbsrv, path_beg, path_dir, path_end, path, path_reg, path_sub, quic_enabled, traffic_is_http, traffic_is_ssl, sc_bytes_in_rate, sc_bytes_out_rate, sc_clr_gpc, sc_clr_gpc0, sc_clr_gpc1, sc0_clr_gpc0, sc0_clr_gpc1, sc1_clr_gpc, sc1_clr_gpc0, sc1_clr_gpc1, sc2_clr_gpc, sc2_clr_gpc0, sc2_clr_gpc1, sc_conn_cnt, sc_conn_cur, sc_conn_rate, sc_get_gpc, sc_get_gpc0, sc_get_gpc1, sc0_get_gpc0, sc0_get_gpc1, sc1_get_gpc0, sc1_get_gpc1, sc2_get_gpc0, sc2_get_gpc1, sc_get_gpt, sc_get_gpt0, sc0_get_gpt0, sc1_get_gpt0, sc2_get_gpt0, sc_glitch_cnt, sc_glitch_rate, sc_gpc_rate, sc_gpc0_rate, sc_gpc1_rate, sc0_gpc0_rate, sc0_gpc1_rate, sc1_gpc0_rate, sc1_gpc1_rate, sc2_gpc0_rate, sc2_gpc1_rate, sc_http_err_cnt, sc_http_err_rate, sc_http_fail_cnt, sc_http_fail_rate, sc_http_req_cnt, sc_http_req_rate, sc_inc_gpc, sc_inc_gpc0, sc_inc_gpc1, sc0_inc_gpc0, sc0_inc_gpc1, sc1_inc_gpc0, sc1_inc_gpc1, sc2_inc_gpc0, sc2_inc_gpc1, sc_sess_cnt, sc_sess_rate, src, src_bytes_in_rate, src_bytes_out_rate, src_clr_gpc, src_clr_gpc0, src_clr_gpc1, src_conn_cnt, src_conn_cur, src_conn_rate, src_get_gpc, src_get_gpc0, src_get_gpc1, src_get_gpt, src_glitch_cnt, src_glitch_rate, src_gpc_rate, src_gpc0_rate, src_gpc1_rate, src_http_err_cnt, src_http_err_rate, src_http_fail_cnt, src_http_fail_rate, src_http_req_cnt, src_http_req_rate, src_inc_gpc, src_inc_gpc0, src_inc_gpc1, src_is_local, src_kbytes_in, src_kbytes_out, src_port, src_sess_cnt, src_sess_rate, ssl_c_ca_commonname, ssl_c_verify_code, ssl_c_verify, ssl_fc_sni, ssl_fc, ssl_hello_type, ssl_sni_beg, ssl_sni_end, ssl_sni_reg, ssl_sni, ssl_sni_sub, stopping, url_param, var, wait_end, custom_acl.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("cust_hdr_beg", "cust_hdr_end", "cust_hdr", "cust_hdr_reg", "cust_hdr_sub", "hdr_beg", "hdr_end", "hdr", "hdr_reg", "hdr_sub", "http_auth", "http_method", "nbsrv", "path_beg", "path_dir", "path_end", "path", "path_reg", "path_sub", "quic_enabled", "traffic_is_http", "traffic_is_ssl", "sc_bytes_in_rate", "sc_bytes_out_rate", "sc_clr_gpc", "sc_clr_gpc0", "sc_clr_gpc1", "sc0_clr_gpc0", "sc0_clr_gpc1", "sc1_clr_gpc", "sc1_clr_gpc0", "sc1_clr_gpc1", "sc2_clr_gpc", "sc2_clr_gpc0", "sc2_clr_gpc1", "sc_conn_cnt", "sc_conn_cur", "sc_conn_rate", "sc_get_gpc", "sc_get_gpc0", "sc_get_gpc1", "sc0_get_gpc0", "sc0_get_gpc1", "sc1_get_gpc0", "sc1_get_gpc1", "sc2_get_gpc0", "sc2_get_gpc1", "sc_get_gpt", "sc_get_gpt0", "sc0_get_gpt0", "sc1_get_gpt0", "sc2_get_gpt0", "sc_glitch_cnt", "sc_glitch_rate", "sc_gpc_rate", "sc_gpc0_rate", "sc_gpc1_rate", "sc0_gpc0_rate", "sc0_gpc1_rate", "sc1_gpc0_rate", "sc1_gpc1_rate", "sc2_gpc0_rate", "sc2_gpc1_rate", "sc_http_err_cnt", "sc_http_err_rate", "sc_http_fail_cnt", "sc_http_fail_rate", "sc_http_req_cnt", "sc_http_req_rate", "sc_inc_gpc", "sc_inc_gpc0", "sc_inc_gpc1", "sc0_inc_gpc0", "sc0_inc_gpc1", "sc1_inc_gpc0", "sc1_inc_gpc1", "sc2_inc_gpc0", "sc2_inc_gpc1", "sc_sess_cnt", "sc_sess_rate", "src", "src_bytes_in_rate", "src_bytes_out_rate", "src_clr_gpc", "src_clr_gpc0", "src_clr_gpc1", "src_conn_cnt", "src_conn_cur", "src_conn_rate", "src_get_gpc", "src_get_gpc0", "src_get_gpc1", "src_get_gpt", "src_glitch_cnt", "src_glitch_rate", "src_gpc_rate", "src_gpc0_rate", "src_gpc1_rate", "src_http_err_cnt", "src_http_err_rate", "src_http_fail_cnt", "src_http_fail_rate", "src_http_req_cnt", "src_http_req_rate", "src_inc_gpc", "src_inc_gpc0", "src_inc_gpc1", "src_is_local", "src_kbytes_in", "src_kbytes_out", "src_port", "src_sess_cnt", "src_sess_rate", "ssl_c_ca_commonname", "ssl_c_verify_code", "ssl_c_verify", "ssl_fc_sni", "ssl_fc", "ssl_hello_type", "ssl_sni_beg", "ssl_sni_end", "ssl_sni_reg", "ssl_sni", "ssl_sni_sub", "stopping", "url_param", "var", "wait_end", "custom_acl")},
			},
			"negate": schema.BoolAttribute{
				MarkdownDescription: "Enable the negate option for this acl.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"case_sensitive": schema.BoolAttribute{
				MarkdownDescription: "Enable the case_sensitive option for this acl.",
				Optional: true, Computed: true,
				Default: booldefault.StaticBool(false),
			},
			"hdr_beg": schema.StringAttribute{
				MarkdownDescription: "ACL hdr_beg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"hdr_end": schema.StringAttribute{
				MarkdownDescription: "ACL hdr_end value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"hdr": schema.StringAttribute{
				MarkdownDescription: "ACL hdr value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"hdr_reg": schema.StringAttribute{
				MarkdownDescription: "ACL hdr_reg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"hdr_sub": schema.StringAttribute{
				MarkdownDescription: "ACL hdr_sub value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"path_beg": schema.StringAttribute{
				MarkdownDescription: "ACL path_beg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"path_end": schema.StringAttribute{
				MarkdownDescription: "ACL path_end value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"path": schema.StringAttribute{
				MarkdownDescription: "ACL path value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"path_reg": schema.StringAttribute{
				MarkdownDescription: "ACL path_reg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"path_dir": schema.StringAttribute{
				MarkdownDescription: "ACL path_dir value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"path_sub": schema.StringAttribute{
				MarkdownDescription: "ACL path_sub value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_beg_name": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_beg_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_beg": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_beg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_end_name": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_end_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_end": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_end value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_name": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_reg_name": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_reg_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_reg": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_reg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_sub_name": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_sub_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"cust_hdr_sub": schema.StringAttribute{
				MarkdownDescription: "ACL cust_hdr_sub value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"url_param": schema.StringAttribute{
				MarkdownDescription: "ACL url_param value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"url_param_value": schema.StringAttribute{
				MarkdownDescription: "ACL url_param_value value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"var": schema.StringAttribute{
				MarkdownDescription: "ACL var value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"var_value": schema.StringAttribute{
				MarkdownDescription: "ACL var_value value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"var_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL var_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"ssl_c_verify_code": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_c_verify_code value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_c_ca_commonname": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_c_ca_commonname value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_hello_type": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_hello_type option. One of: , x0, x1, x2.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "x0", "x1", "x2")},
			},
			"src": schema.StringAttribute{
				MarkdownDescription: "ACL src value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_bytes_in_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_bytes_in_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_bytes_in_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_bytes_in_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_bytes_out_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_bytes_out_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_bytes_out_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_bytes_out_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_conn_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_conn_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_conn_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL src_conn_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_conn_cur_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_conn_cur_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_conn_cur": schema.StringAttribute{
				MarkdownDescription: "ACL src_conn_cur value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_conn_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_conn_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_conn_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_conn_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_http_err_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_err_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_err_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_err_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_http_err_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_err_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_err_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_err_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_http_req_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_req_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_req_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_req_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_http_req_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_req_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_req_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_req_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_kbytes_in_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_kbytes_in_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_kbytes_in": schema.StringAttribute{
				MarkdownDescription: "ACL src_kbytes_in value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_kbytes_out_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_kbytes_out_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_kbytes_out": schema.StringAttribute{
				MarkdownDescription: "ACL src_kbytes_out value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_port_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_port_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_port": schema.StringAttribute{
				MarkdownDescription: "ACL src_port value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_sess_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_sess_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_sess_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL src_sess_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_sess_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_sess_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_sess_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_sess_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"nbsrv": schema.StringAttribute{
				MarkdownDescription: "ACL nbsrv value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"nbsrv_backend": schema.StringAttribute{
				MarkdownDescription: "ACL nbsrv_backend (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_fc_sni": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_fc_sni value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_sni": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_sni value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_sni_sub": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_sni_sub value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_sni_beg": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_sni_beg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_sni_end": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_sni_end value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"ssl_sni_reg": schema.StringAttribute{
				MarkdownDescription: "ACL ssl_sni_reg value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"custom_acl": schema.StringAttribute{
				MarkdownDescription: "ACL custom_acl value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"value": schema.StringAttribute{
				MarkdownDescription: "ACL value value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"urlparam": schema.StringAttribute{
				MarkdownDescription: "ACL urlparam value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"query_backend": schema.StringAttribute{
				MarkdownDescription: "ACL query_backend (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"allowed_users": schema.SetAttribute{
				MarkdownDescription: "List of allowed_users values for this acl.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"allowed_groups": schema.SetAttribute{
				MarkdownDescription: "List of allowed_groups values for this acl.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"http_method": schema.SetAttribute{
				MarkdownDescription: "Selected http_method values for this acl. One or more of: CONNECT, DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT, TRACE.",
				ElementType: types.StringType,
				Optional: true, Computed: true,
				Default: setdefault.StaticValue(tools.EmptySetValue(types.StringType)),
			},
			"sc_bytes_in_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_bytes_in_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_bytes_in_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_bytes_in_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_bytes_out_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_bytes_out_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_bytes_out_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_bytes_out_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_clr_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_clr_gpc_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_clr_gpc": schema.StringAttribute{
				MarkdownDescription: "ACL sc_clr_gpc value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_conn_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_conn_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_conn_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL sc_conn_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_conn_cur_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_conn_cur_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_conn_cur": schema.StringAttribute{
				MarkdownDescription: "ACL sc_conn_cur value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_conn_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_conn_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_conn_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_conn_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_get_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpc_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpc": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpc value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_glitch_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_glitch_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_glitch_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL sc_glitch_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_glitch_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_glitch_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_glitch_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_glitch_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_gpc_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_gpc_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_gpc_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_gpc_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_http_err_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_err_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_err_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_err_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_http_err_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_err_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_err_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_err_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_http_fail_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_fail_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_fail_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_fail_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_http_fail_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_fail_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_fail_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_fail_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_http_req_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_req_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_req_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_req_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_http_req_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_req_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_http_req_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_http_req_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_inc_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_inc_gpc_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_inc_gpc": schema.StringAttribute{
				MarkdownDescription: "ACL sc_inc_gpc value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_sess_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_sess_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_sess_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL sc_sess_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_sess_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_sess_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_sess_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_sess_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_get_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpc_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpc": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpc value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_get_gpt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpt": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_glitch_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_glitch_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_glitch_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL src_glitch_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_glitch_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_glitch_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_glitch_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_glitch_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_gpc_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_gpc_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_gpc_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_gpc_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_http_fail_cnt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_fail_cnt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_fail_cnt": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_fail_cnt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_http_fail_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_fail_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_http_fail_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_http_fail_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_inc_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_inc_gpc_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_inc_gpc": schema.StringAttribute{
				MarkdownDescription: "ACL src_inc_gpc value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_clr_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc_clr_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_clr_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc_clr_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_clr_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_clr_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_clr_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_clr_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_clr_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_clr_gpc_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_clr_gpc": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_clr_gpc value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_clr_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_clr_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_clr_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_clr_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_clr_gpc_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_clr_gpc_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_clr_gpc": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_clr_gpc value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_clr_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_clr_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_clr_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_clr_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_get_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_get_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_get_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_get_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_get_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_get_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_get_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_get_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_get_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_get_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_get_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_get_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_get_gpt_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpt_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpt": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpt value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_get_gpt0": schema.StringAttribute{
				MarkdownDescription: "ACL sc_get_gpt0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_get_gpt0": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_get_gpt0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_get_gpt0": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_get_gpt0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_get_gpt0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_get_gpt0": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_get_gpt0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_gpc0_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_gpc0_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_gpc1_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc_gpc1_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_gpc0_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_gpc0_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_gpc1_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_gpc1_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_gpc0_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_gpc0_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_gpc1_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_gpc1_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_gpc0_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_gpc0_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_gpc1_rate": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_gpc1_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_inc_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc_inc_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc_inc_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc_inc_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_inc_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_inc_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc0_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc0_inc_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc0_inc_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_inc_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_inc_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc1_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc1_inc_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc1_inc_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_inc_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_inc_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc2_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"sc2_inc_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL sc2_inc_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_clr_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_clr_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL src_clr_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_clr_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_clr_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL src_clr_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_get_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_get_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_get_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL src_get_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_gpc0_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_gpc0_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_gpc0_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_gpc1_rate_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_gpc1_rate": schema.StringAttribute{
				MarkdownDescription: "ACL src_gpc1_rate value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_inc_gpc0_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_inc_gpc0": schema.StringAttribute{
				MarkdownDescription: "ACL src_inc_gpc0 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"src_inc_gpc1_comparison": schema.StringAttribute{
				MarkdownDescription: "ACL src_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
				Validators: []validator.String{stringvalidator.OneOf("", "gt", "ge", "eq", "lt", "le")},
			},
			"src_inc_gpc1": schema.StringAttribute{
				MarkdownDescription: "ACL src_inc_gpc1 value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"gpc_number": schema.StringAttribute{
				MarkdownDescription: "ACL gpc_number value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"gpt_number": schema.StringAttribute{
				MarkdownDescription: "ACL gpt_number value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"sc_number": schema.StringAttribute{
				MarkdownDescription: "ACL sc_number value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"table_name": schema.StringAttribute{
				MarkdownDescription: "ACL table_name value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"mapfile": schema.StringAttribute{
				MarkdownDescription: "ACL mapfile (UUID reference).",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
			"converter": schema.StringAttribute{
				MarkdownDescription: "ACL converter value.",
				Optional: true, Computed: true,
				Default: stringdefault.StaticString(""),
			},
		},
	}
}

func aclDataSourceSchema() dschema.Schema {
	return dschema.Schema{
		MarkdownDescription: "Read an OPNsense HAProxy condition/ACL.",
		Attributes: map[string]dschema.Attribute{
			"id": dschema.StringAttribute{
				MarkdownDescription: "UUID of the haproxy acl.",
				Required: true,
				Validators: []validator.String{stringvalidator.LengthAtLeast(1)},
			},
			"internal_id": dschema.StringAttribute{MarkdownDescription: "ACL internal_id value.", Computed: true},
			"name": dschema.StringAttribute{MarkdownDescription: "ACL name value.", Computed: true},
			"description": dschema.StringAttribute{MarkdownDescription: "ACL description value.", Computed: true},
			"expression": dschema.StringAttribute{MarkdownDescription: "ACL expression option. One of: cust_hdr_beg, cust_hdr_end, cust_hdr, cust_hdr_reg, cust_hdr_sub, hdr_beg, hdr_end, hdr, hdr_reg, hdr_sub, http_auth, http_method, nbsrv, path_beg, path_dir, path_end, path, path_reg, path_sub, quic_enabled, traffic_is_http, traffic_is_ssl, sc_bytes_in_rate, sc_bytes_out_rate, sc_clr_gpc, sc_clr_gpc0, sc_clr_gpc1, sc0_clr_gpc0, sc0_clr_gpc1, sc1_clr_gpc, sc1_clr_gpc0, sc1_clr_gpc1, sc2_clr_gpc, sc2_clr_gpc0, sc2_clr_gpc1, sc_conn_cnt, sc_conn_cur, sc_conn_rate, sc_get_gpc, sc_get_gpc0, sc_get_gpc1, sc0_get_gpc0, sc0_get_gpc1, sc1_get_gpc0, sc1_get_gpc1, sc2_get_gpc0, sc2_get_gpc1, sc_get_gpt, sc_get_gpt0, sc0_get_gpt0, sc1_get_gpt0, sc2_get_gpt0, sc_glitch_cnt, sc_glitch_rate, sc_gpc_rate, sc_gpc0_rate, sc_gpc1_rate, sc0_gpc0_rate, sc0_gpc1_rate, sc1_gpc0_rate, sc1_gpc1_rate, sc2_gpc0_rate, sc2_gpc1_rate, sc_http_err_cnt, sc_http_err_rate, sc_http_fail_cnt, sc_http_fail_rate, sc_http_req_cnt, sc_http_req_rate, sc_inc_gpc, sc_inc_gpc0, sc_inc_gpc1, sc0_inc_gpc0, sc0_inc_gpc1, sc1_inc_gpc0, sc1_inc_gpc1, sc2_inc_gpc0, sc2_inc_gpc1, sc_sess_cnt, sc_sess_rate, src, src_bytes_in_rate, src_bytes_out_rate, src_clr_gpc, src_clr_gpc0, src_clr_gpc1, src_conn_cnt, src_conn_cur, src_conn_rate, src_get_gpc, src_get_gpc0, src_get_gpc1, src_get_gpt, src_glitch_cnt, src_glitch_rate, src_gpc_rate, src_gpc0_rate, src_gpc1_rate, src_http_err_cnt, src_http_err_rate, src_http_fail_cnt, src_http_fail_rate, src_http_req_cnt, src_http_req_rate, src_inc_gpc, src_inc_gpc0, src_inc_gpc1, src_is_local, src_kbytes_in, src_kbytes_out, src_port, src_sess_cnt, src_sess_rate, ssl_c_ca_commonname, ssl_c_verify_code, ssl_c_verify, ssl_fc_sni, ssl_fc, ssl_hello_type, ssl_sni_beg, ssl_sni_end, ssl_sni_reg, ssl_sni, ssl_sni_sub, stopping, url_param, var, wait_end, custom_acl.", Computed: true},
			"negate": dschema.BoolAttribute{MarkdownDescription: "Enable the negate option for this acl.", Computed: true},
			"case_sensitive": dschema.BoolAttribute{MarkdownDescription: "Enable the case_sensitive option for this acl.", Computed: true},
			"hdr_beg": dschema.StringAttribute{MarkdownDescription: "ACL hdr_beg value.", Computed: true},
			"hdr_end": dschema.StringAttribute{MarkdownDescription: "ACL hdr_end value.", Computed: true},
			"hdr": dschema.StringAttribute{MarkdownDescription: "ACL hdr value.", Computed: true},
			"hdr_reg": dschema.StringAttribute{MarkdownDescription: "ACL hdr_reg value.", Computed: true},
			"hdr_sub": dschema.StringAttribute{MarkdownDescription: "ACL hdr_sub value.", Computed: true},
			"path_beg": dschema.StringAttribute{MarkdownDescription: "ACL path_beg value.", Computed: true},
			"path_end": dschema.StringAttribute{MarkdownDescription: "ACL path_end value.", Computed: true},
			"path": dschema.StringAttribute{MarkdownDescription: "ACL path value.", Computed: true},
			"path_reg": dschema.StringAttribute{MarkdownDescription: "ACL path_reg value.", Computed: true},
			"path_dir": dschema.StringAttribute{MarkdownDescription: "ACL path_dir value.", Computed: true},
			"path_sub": dschema.StringAttribute{MarkdownDescription: "ACL path_sub value.", Computed: true},
			"cust_hdr_beg_name": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_beg_name value.", Computed: true},
			"cust_hdr_beg": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_beg value.", Computed: true},
			"cust_hdr_end_name": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_end_name value.", Computed: true},
			"cust_hdr_end": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_end value.", Computed: true},
			"cust_hdr_name": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_name value.", Computed: true},
			"cust_hdr": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr value.", Computed: true},
			"cust_hdr_reg_name": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_reg_name value.", Computed: true},
			"cust_hdr_reg": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_reg value.", Computed: true},
			"cust_hdr_sub_name": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_sub_name value.", Computed: true},
			"cust_hdr_sub": dschema.StringAttribute{MarkdownDescription: "ACL cust_hdr_sub value.", Computed: true},
			"url_param": dschema.StringAttribute{MarkdownDescription: "ACL url_param value.", Computed: true},
			"url_param_value": dschema.StringAttribute{MarkdownDescription: "ACL url_param_value value.", Computed: true},
			"var": dschema.StringAttribute{MarkdownDescription: "ACL var value.", Computed: true},
			"var_value": dschema.StringAttribute{MarkdownDescription: "ACL var_value value.", Computed: true},
			"var_comparison": dschema.StringAttribute{MarkdownDescription: "ACL var_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"ssl_c_verify_code": dschema.StringAttribute{MarkdownDescription: "ACL ssl_c_verify_code value.", Computed: true},
			"ssl_c_ca_commonname": dschema.StringAttribute{MarkdownDescription: "ACL ssl_c_ca_commonname value.", Computed: true},
			"ssl_hello_type": dschema.StringAttribute{MarkdownDescription: "ACL ssl_hello_type option. One of: , x0, x1, x2.", Computed: true},
			"src": dschema.StringAttribute{MarkdownDescription: "ACL src value.", Computed: true},
			"src_bytes_in_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_bytes_in_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_bytes_in_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_bytes_in_rate value.", Computed: true},
			"src_bytes_out_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_bytes_out_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_bytes_out_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_bytes_out_rate value.", Computed: true},
			"src_conn_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_conn_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_conn_cnt": dschema.StringAttribute{MarkdownDescription: "ACL src_conn_cnt value.", Computed: true},
			"src_conn_cur_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_conn_cur_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_conn_cur": dschema.StringAttribute{MarkdownDescription: "ACL src_conn_cur value.", Computed: true},
			"src_conn_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_conn_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_conn_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_conn_rate value.", Computed: true},
			"src_http_err_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_http_err_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_http_err_cnt": dschema.StringAttribute{MarkdownDescription: "ACL src_http_err_cnt value.", Computed: true},
			"src_http_err_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_http_err_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_http_err_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_http_err_rate value.", Computed: true},
			"src_http_req_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_http_req_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_http_req_cnt": dschema.StringAttribute{MarkdownDescription: "ACL src_http_req_cnt value.", Computed: true},
			"src_http_req_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_http_req_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_http_req_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_http_req_rate value.", Computed: true},
			"src_kbytes_in_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_kbytes_in_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_kbytes_in": dschema.StringAttribute{MarkdownDescription: "ACL src_kbytes_in value.", Computed: true},
			"src_kbytes_out_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_kbytes_out_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_kbytes_out": dschema.StringAttribute{MarkdownDescription: "ACL src_kbytes_out value.", Computed: true},
			"src_port_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_port_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_port": dschema.StringAttribute{MarkdownDescription: "ACL src_port value.", Computed: true},
			"src_sess_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_sess_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_sess_cnt": dschema.StringAttribute{MarkdownDescription: "ACL src_sess_cnt value.", Computed: true},
			"src_sess_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_sess_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_sess_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_sess_rate value.", Computed: true},
			"nbsrv": dschema.StringAttribute{MarkdownDescription: "ACL nbsrv value.", Computed: true},
			"nbsrv_backend": dschema.StringAttribute{MarkdownDescription: "ACL nbsrv_backend (UUID reference).", Computed: true},
			"ssl_fc_sni": dschema.StringAttribute{MarkdownDescription: "ACL ssl_fc_sni value.", Computed: true},
			"ssl_sni": dschema.StringAttribute{MarkdownDescription: "ACL ssl_sni value.", Computed: true},
			"ssl_sni_sub": dschema.StringAttribute{MarkdownDescription: "ACL ssl_sni_sub value.", Computed: true},
			"ssl_sni_beg": dschema.StringAttribute{MarkdownDescription: "ACL ssl_sni_beg value.", Computed: true},
			"ssl_sni_end": dschema.StringAttribute{MarkdownDescription: "ACL ssl_sni_end value.", Computed: true},
			"ssl_sni_reg": dschema.StringAttribute{MarkdownDescription: "ACL ssl_sni_reg value.", Computed: true},
			"custom_acl": dschema.StringAttribute{MarkdownDescription: "ACL custom_acl value.", Computed: true},
			"value": dschema.StringAttribute{MarkdownDescription: "ACL value value.", Computed: true},
			"urlparam": dschema.StringAttribute{MarkdownDescription: "ACL urlparam value.", Computed: true},
			"query_backend": dschema.StringAttribute{MarkdownDescription: "ACL query_backend (UUID reference).", Computed: true},
			"allowed_users": dschema.SetAttribute{MarkdownDescription: "List of allowed_users values for this acl.", ElementType: types.StringType, Computed: true},
			"allowed_groups": dschema.SetAttribute{MarkdownDescription: "List of allowed_groups values for this acl.", ElementType: types.StringType, Computed: true},
			"http_method": dschema.SetAttribute{MarkdownDescription: "Selected http_method values for this acl. One or more of: CONNECT, DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT, TRACE.", ElementType: types.StringType, Computed: true},
			"sc_bytes_in_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_bytes_in_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_bytes_in_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_bytes_in_rate value.", Computed: true},
			"sc_bytes_out_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_bytes_out_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_bytes_out_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_bytes_out_rate value.", Computed: true},
			"sc_clr_gpc_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_clr_gpc_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_clr_gpc": dschema.StringAttribute{MarkdownDescription: "ACL sc_clr_gpc value.", Computed: true},
			"sc_conn_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_conn_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_conn_cnt": dschema.StringAttribute{MarkdownDescription: "ACL sc_conn_cnt value.", Computed: true},
			"sc_conn_cur_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_conn_cur_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_conn_cur": dschema.StringAttribute{MarkdownDescription: "ACL sc_conn_cur value.", Computed: true},
			"sc_conn_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_conn_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_conn_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_conn_rate value.", Computed: true},
			"sc_get_gpc_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpc_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_get_gpc": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpc value.", Computed: true},
			"sc_glitch_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_glitch_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_glitch_cnt": dschema.StringAttribute{MarkdownDescription: "ACL sc_glitch_cnt value.", Computed: true},
			"sc_glitch_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_glitch_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_glitch_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_glitch_rate value.", Computed: true},
			"sc_gpc_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_gpc_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_gpc_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_gpc_rate value.", Computed: true},
			"sc_http_err_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_err_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_http_err_cnt": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_err_cnt value.", Computed: true},
			"sc_http_err_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_err_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_http_err_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_err_rate value.", Computed: true},
			"sc_http_fail_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_fail_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_http_fail_cnt": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_fail_cnt value.", Computed: true},
			"sc_http_fail_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_fail_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_http_fail_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_fail_rate value.", Computed: true},
			"sc_http_req_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_req_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_http_req_cnt": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_req_cnt value.", Computed: true},
			"sc_http_req_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_req_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_http_req_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_http_req_rate value.", Computed: true},
			"sc_inc_gpc_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_inc_gpc_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_inc_gpc": dschema.StringAttribute{MarkdownDescription: "ACL sc_inc_gpc value.", Computed: true},
			"sc_sess_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_sess_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_sess_cnt": dschema.StringAttribute{MarkdownDescription: "ACL sc_sess_cnt value.", Computed: true},
			"sc_sess_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_sess_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_sess_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_sess_rate value.", Computed: true},
			"src_get_gpc_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpc_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_get_gpc": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpc value.", Computed: true},
			"src_get_gpt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_get_gpt": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpt value.", Computed: true},
			"src_glitch_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_glitch_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_glitch_cnt": dschema.StringAttribute{MarkdownDescription: "ACL src_glitch_cnt value.", Computed: true},
			"src_glitch_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_glitch_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_glitch_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_glitch_rate value.", Computed: true},
			"src_gpc_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_gpc_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_gpc_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_gpc_rate value.", Computed: true},
			"src_http_fail_cnt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_http_fail_cnt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_http_fail_cnt": dschema.StringAttribute{MarkdownDescription: "ACL src_http_fail_cnt value.", Computed: true},
			"src_http_fail_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_http_fail_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_http_fail_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_http_fail_rate value.", Computed: true},
			"src_inc_gpc_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_inc_gpc_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_inc_gpc": dschema.StringAttribute{MarkdownDescription: "ACL src_inc_gpc value.", Computed: true},
			"sc_clr_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_clr_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc_clr_gpc0 value.", Computed: true},
			"sc_clr_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_clr_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc_clr_gpc1 value.", Computed: true},
			"sc0_clr_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_clr_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc0_clr_gpc0 value.", Computed: true},
			"sc0_clr_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_clr_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc0_clr_gpc1 value.", Computed: true},
			"sc1_clr_gpc_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_clr_gpc_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_clr_gpc": dschema.StringAttribute{MarkdownDescription: "ACL sc1_clr_gpc value.", Computed: true},
			"sc1_clr_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_clr_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc1_clr_gpc0 value.", Computed: true},
			"sc1_clr_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_clr_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc1_clr_gpc1 value.", Computed: true},
			"sc2_clr_gpc_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_clr_gpc_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_clr_gpc": dschema.StringAttribute{MarkdownDescription: "ACL sc2_clr_gpc value.", Computed: true},
			"sc2_clr_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_clr_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc2_clr_gpc0 value.", Computed: true},
			"sc2_clr_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_clr_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc2_clr_gpc1 value.", Computed: true},
			"sc_get_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_get_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpc0 value.", Computed: true},
			"sc_get_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_get_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpc1 value.", Computed: true},
			"sc0_get_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_get_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc0_get_gpc0 value.", Computed: true},
			"sc0_get_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_get_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc0_get_gpc1 value.", Computed: true},
			"sc1_get_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_get_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc1_get_gpc0 value.", Computed: true},
			"sc1_get_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_get_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc1_get_gpc1 value.", Computed: true},
			"sc2_get_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_get_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc2_get_gpc0 value.", Computed: true},
			"sc2_get_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_get_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc2_get_gpc1 value.", Computed: true},
			"sc_get_gpt_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpt_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_get_gpt": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpt value.", Computed: true},
			"sc_get_gpt0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_get_gpt0": dschema.StringAttribute{MarkdownDescription: "ACL sc_get_gpt0 value.", Computed: true},
			"sc0_get_gpt0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_get_gpt0": dschema.StringAttribute{MarkdownDescription: "ACL sc0_get_gpt0 value.", Computed: true},
			"sc1_get_gpt0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_get_gpt0": dschema.StringAttribute{MarkdownDescription: "ACL sc1_get_gpt0 value.", Computed: true},
			"sc2_get_gpt0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_get_gpt0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_get_gpt0": dschema.StringAttribute{MarkdownDescription: "ACL sc2_get_gpt0 value.", Computed: true},
			"sc_gpc0_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_gpc0_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_gpc0_rate value.", Computed: true},
			"sc_gpc1_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_gpc1_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc_gpc1_rate value.", Computed: true},
			"sc0_gpc0_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_gpc0_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc0_gpc0_rate value.", Computed: true},
			"sc0_gpc1_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_gpc1_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc0_gpc1_rate value.", Computed: true},
			"sc1_gpc0_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_gpc0_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc1_gpc0_rate value.", Computed: true},
			"sc1_gpc1_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_gpc1_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc1_gpc1_rate value.", Computed: true},
			"sc2_gpc0_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_gpc0_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc2_gpc0_rate value.", Computed: true},
			"sc2_gpc1_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_gpc1_rate": dschema.StringAttribute{MarkdownDescription: "ACL sc2_gpc1_rate value.", Computed: true},
			"sc_inc_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_inc_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc_inc_gpc0 value.", Computed: true},
			"sc_inc_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc_inc_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc_inc_gpc1 value.", Computed: true},
			"sc0_inc_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_inc_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc0_inc_gpc0 value.", Computed: true},
			"sc0_inc_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc0_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc0_inc_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc0_inc_gpc1 value.", Computed: true},
			"sc1_inc_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_inc_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc1_inc_gpc0 value.", Computed: true},
			"sc1_inc_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc1_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc1_inc_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc1_inc_gpc1 value.", Computed: true},
			"sc2_inc_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_inc_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL sc2_inc_gpc0 value.", Computed: true},
			"sc2_inc_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL sc2_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"sc2_inc_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL sc2_inc_gpc1 value.", Computed: true},
			"src_clr_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_clr_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_clr_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL src_clr_gpc0 value.", Computed: true},
			"src_clr_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_clr_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_clr_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL src_clr_gpc1 value.", Computed: true},
			"src_get_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_get_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpc0 value.", Computed: true},
			"src_get_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_get_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL src_get_gpc1 value.", Computed: true},
			"src_gpc0_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_gpc0_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_gpc0_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_gpc0_rate value.", Computed: true},
			"src_gpc1_rate_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_gpc1_rate_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_gpc1_rate": dschema.StringAttribute{MarkdownDescription: "ACL src_gpc1_rate value.", Computed: true},
			"src_inc_gpc0_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_inc_gpc0_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_inc_gpc0": dschema.StringAttribute{MarkdownDescription: "ACL src_inc_gpc0 value.", Computed: true},
			"src_inc_gpc1_comparison": dschema.StringAttribute{MarkdownDescription: "ACL src_inc_gpc1_comparison option. One of: , gt, ge, eq, lt, le.", Computed: true},
			"src_inc_gpc1": dschema.StringAttribute{MarkdownDescription: "ACL src_inc_gpc1 value.", Computed: true},
			"gpc_number": dschema.StringAttribute{MarkdownDescription: "ACL gpc_number value.", Computed: true},
			"gpt_number": dschema.StringAttribute{MarkdownDescription: "ACL gpt_number value.", Computed: true},
			"sc_number": dschema.StringAttribute{MarkdownDescription: "ACL sc_number value.", Computed: true},
			"table_name": dschema.StringAttribute{MarkdownDescription: "ACL table_name value.", Computed: true},
			"mapfile": dschema.StringAttribute{MarkdownDescription: "ACL mapfile (UUID reference).", Computed: true},
			"converter": dschema.StringAttribute{MarkdownDescription: "ACL converter value.", Computed: true},
		},
	}
}
