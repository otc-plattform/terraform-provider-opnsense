---
page_title: "opnsense_haproxy_resolver Data Source - terraform-provider-opnsense"
description: Read an OPNsense HAProxy DNS resolver.
---

# opnsense_haproxy_resolver (Data Source)

```terraform
data "opnsense_haproxy_resolver" "example" {
  id = "<resolver-uuid>"
}
```

## Schema

### Required

- `id` (String) Resolver UUID.

### Read-Only

- `name`, `description` (String) Resolver name and description.
- `enabled`, `parse_resolv_conf` (Boolean) Resolver settings.
- `nameservers` (Set of String) DNS server addresses.
- `resolve_retries`, `timeout_resolve`, `timeout_retry`, `accepted_payload_size` (String) DNS resolver settings.
- `hold_valid`, `hold_obsolete`, `hold_refused`, `hold_nx`, `hold_timeout`, `hold_other` (String) DNS response hold times.
