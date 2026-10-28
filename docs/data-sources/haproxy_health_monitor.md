---
page_title: "opnsense_haproxy_health_monitor Data Source - terraform-provider-opnsense"
description: Read an OPNsense HAProxy health monitor.
---

# opnsense_haproxy_health_monitor (Data Source)

```terraform
data "opnsense_haproxy_health_monitor" "example" {
  id = "<health-monitor-uuid>"
}
```

## Schema

### Required

- `id` (String) Health monitor UUID.

### Read-Only

- `name`, `description`, `type`, `interval` (String) Monitor details.
- `ssl`, `ssl_sni`, `check_port` (String) TLS and port settings.
- `http_method`, `http_uri`, `http_version`, `http_host`, `http_expression`, `http_value` (String) HTTP check settings.
- `http_expression_enabled`, `http_negate` (Boolean) HTTP response matching settings.
- `tcp_enabled`, `tcp_negate` (Boolean) TCP response matching settings.
- `tcp_send_value`, `tcp_match_type`, `tcp_match_value` (String) TCP send and match settings.
- `agent_port`, `mysql_user`, `postgres_user`, `smtp_domain`, `esmtp_domain` (String) Protocol-specific settings.
- `mysql_post41` (Boolean) Whether MySQL 4.1 or newer checks are enabled.
