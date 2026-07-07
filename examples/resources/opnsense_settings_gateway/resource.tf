resource "opnsense_settings_gateway" "test" {
  name            = "TEST_123"
  interface       = "wan"
  ip_protocol     = "inet"
  description     = "Terraform managed gateway"
  priority        = 0
  weight          = 1
  monitor_disable = false
  disabled        = false
}
