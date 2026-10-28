---
page_title: "opnsense_haproxy_health_monitor Resource - terraform-provider-opnsense"
description: Manage an OPNsense HAProxy health monitor.
---

# opnsense_haproxy_health_monitor (Resource)

Manages a health monitor under HAProxy Settings, Rules & Checks, Health Monitors. Reference its `id` from an HAProxy backend's `health_check` attribute and set `health_check_enabled` on the backend.

```terraform
resource "opnsense_haproxy_health_monitor" "http" {
  name                    = "http-health"
  type                    = "http"
  interval                = "5s"
  http_method             = "get"
  http_uri                = "/health"
  http_expression_enabled = true
  http_expression         = "status"
  http_value              = "200"
}

resource "opnsense_haproxy_apply" "configuration" {
  triggers = {
    health_monitor = sha256(jsonencode(opnsense_haproxy_health_monitor.http))
    # Include all other relevant HAProxy resources as well.
  }

  depends_on = [opnsense_haproxy_health_monitor.http]
}
```

Use the existing `opnsense_haproxy_apply` resource for the whole HAProxy configuration. Include every relevant HAProxy resource in its triggers and dependencies. Changes are saved but not activated until HAProxy is reconfigured.

## Schema

### Required

- `name` (String) Health monitor name.

### Optional

- `description` (String) Description.
- `type` (String) Check type: `http`, `tcp`, `agent`, `ldap`, `mysql`, `pgsql`, `redis`, `smtp`, `esmtp` or `ssl`. Defaults to `http`.
- `interval` (String) Check interval. Defaults to `2s`.
- `ssl` (String) TLS mode: `nopref`, `ssl`, `sslsni` or `nossl`. Defaults to `nopref`.
- `ssl_sni` (String) SNI name for TLS checks.
- `check_port` (String) Override check port.
- `http_method`, `http_uri`, `http_version`, `http_host` (String) HTTP check settings.
- `http_expression_enabled`, `http_negate` (Boolean) Enable or negate an HTTP response expression.
- `http_expression`, `http_value` (String) HTTP response match type and value.
- `tcp_enabled`, `tcp_negate` (Boolean) Enable TCP send and response matching or negate the match.
- `tcp_send_value`, `tcp_match_type`, `tcp_match_value` (String) TCP send and match settings.
- `agent_port` (String) Agent check port.
- `mysql_user` (String), `mysql_post41` (Boolean), `postgres_user` (String) Database health check settings.
- `smtp_domain`, `esmtp_domain` (String) Mail check domain names.

### Read-Only

- `id` (String) Health monitor UUID.

## Import

```shell
terraform import opnsense_haproxy_health_monitor.http <health-monitor-uuid>
```
