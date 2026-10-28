---
page_title: "opnsense_haproxy_apply Resource - terraform-provider-opnsense"
description: Validate and activate the staged HAProxy configuration.
---

# opnsense_haproxy_apply (Resource)

Runs the HAProxy syntax check and, if it succeeds, reconfigures the service. Critical errors fail the Terraform apply and include the HAProxy output; warnings are reported without failing. A missing or unrecognized test result is treated as an error. OPNsense can return an empty result for a valid configuration with no active listeners; that is accepted. Removing this resource from Terraform does not stop or change HAProxy.

**Use one apply resource for the complete HAProxy configuration.** It runs when first created and whenever `triggers` changes. It does not monitor or automatically detect changes on the firewall. Include every relevant HAProxy resource in the triggers and dependencies. `depends_on` alone controls ordering but does **not** cause the action to run again when a dependency changes.

```terraform
resource "opnsense_haproxy_apply" "configuration" {
  triggers = {
    settings = sha256(jsonencode(opnsense_haproxy_settings.main))
    backend  = sha256(jsonencode(opnsense_haproxy_backend.application))
    frontend = sha256(jsonencode(opnsense_haproxy_frontend.https))
  }

  # Also include all other HAProxy resources managed by this configuration,
  # such as servers, ACLs, actions, resolvers and health monitors, in triggers and depends_on.
  depends_on = [
    opnsense_haproxy_settings.main,
    opnsense_haproxy_backend.application,
    opnsense_haproxy_frontend.https,
  ]
}
```

If the check fails, objects saved earlier in the same Terraform apply remain saved in OPNsense and in Terraform state. The action does not roll those changes back. Correct the configuration and apply again. Out-of-band changes require changing a trigger or replacing this resource to rerun the action.

## Schema

### Required

- `triggers` (Map of String) Values whose changes trigger another validation and reconfiguration. Trigger values are stored in Terraform state; avoid putting secrets in them directly.

### Read-Only

- `id` (String) Identifier of the apply action.
