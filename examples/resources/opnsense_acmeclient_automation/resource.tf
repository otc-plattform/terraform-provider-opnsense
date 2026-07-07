resource "opnsense_acmeclient_automation" "restart_gui" {
  enabled     = true
  name        = "tf-example-automation"
  description = "Restart GUI after issuance"
  type        = "configd_restart_gui"
}
