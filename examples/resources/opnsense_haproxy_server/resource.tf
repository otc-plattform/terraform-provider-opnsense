resource "opnsense_haproxy_server" "example" {
  enabled              = false
  name                 = "example_server"
  description          = "Example HAProxy real server"
  type                 = "static"
  address              = "192.0.2.10"
  port                 = "8080"
  mode                 = "active"
  multiplexer_protocol = "unspecified"
  ssl                  = false
  ssl_verify           = true
}
