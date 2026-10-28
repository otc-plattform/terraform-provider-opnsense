---
page_title: "opnsense_haproxy_resolver Resource - terraform-provider-opnsense"
description: Manage an OPNsense HAProxy DNS resolver.
---

# opnsense_haproxy_resolver (Resource)

Manages an HAProxy DNS resolver. Reference its `id` from a backend or server's `linked_resolver` attribute.

```terraform
resource "opnsense_haproxy_resolver" "example" {
  name        = "dns"
  nameservers = ["1.1.1.1:53", "8.8.8.8:53"]
}
```

## Schema

### Required

- `name` (String) Resolver name.

### Optional

- `enabled` (Boolean) Defaults to `true`.
- `description` (String) Resolver description.
- `nameservers` (Set of String) DNS server addresses, including ports.
- `parse_resolv_conf` (Boolean) Use the system resolv.conf; defaults to `false`.
- `resolve_retries` (String) Number of DNS retries; defaults to `"3"`.
- `timeout_resolve`, `timeout_retry` (String) DNS timeouts; each defaults to `"1s"`.
- `accepted_payload_size` (String) Maximum accepted DNS payload size.
- `hold_valid`, `hold_obsolete`, `hold_refused`, `hold_nx`, `hold_timeout`, `hold_other` (String) DNS response hold times.

### Read-Only

- `id` (String) Resolver UUID.

## Import

```shell
terraform import opnsense_haproxy_resolver.example <resolver-uuid>
```
