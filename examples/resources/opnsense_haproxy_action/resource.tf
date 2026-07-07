resource "opnsense_haproxy_server" "acme" {
  enabled              = false
  name                 = "example_acme_server"
  type                 = "static"
  address              = "127.0.0.1"
  port                 = "43580"
  mode                 = "active"
  multiplexer_protocol = "unspecified"
  ssl                  = false
  ssl_verify           = true
}

resource "opnsense_haproxy_backend" "acme" {
  enabled                  = false
  name                     = "example_acme_backend"
  mode                     = "http"
  algorithm                = "source"
  random_draws             = "2"
  linked_servers           = [opnsense_haproxy_server.acme.id]
  health_check_enabled     = false
  health_check_proxy_proto = "backend"
  http2_enabled            = true
  ba_advertised_protocols  = ["h2", "http11"]
  persistence              = "sticktable"
  tuning_httpreuse         = "safe"
}

resource "opnsense_haproxy_acl" "acme_challenge" {
  name           = "example_acme_challenge"
  description    = "Match ACME HTTP-01 challenge requests"
  expression     = "path_beg"
  negate         = false
  case_sensitive = false
  path_beg       = "/.well-known/acme-challenge/"
}

resource "opnsense_haproxy_action" "acme_challenge" {
  enabled     = false
  name        = "example_acme_redirect"
  description = "Route ACME challenge requests to the ACME backend"
  test_type   = "if"
  linked_acls = [opnsense_haproxy_acl.acme_challenge.id]
  operator    = "and"
  type        = "use_backend"
  use_backend = opnsense_haproxy_backend.acme.id
}
