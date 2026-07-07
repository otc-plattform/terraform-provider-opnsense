resource "opnsense_acmeclient_settings" "example" {
  enabled             = false
  auto_renewal        = true
  haproxy_integration = false
  log_level           = "debug3"
  show_intro          = false
  challenge_port      = 42580
  tls_challenge_port  = 42581
  restart_timeout     = 600
}
