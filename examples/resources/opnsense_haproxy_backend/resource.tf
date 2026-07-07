resource "opnsense_haproxy_server" "app" {
  enabled              = false
  name                 = "example_app_server"
  description          = "Example backend member"
  type                 = "static"
  address              = "192.0.2.10"
  port                 = "8080"
  mode                 = "active"
  multiplexer_protocol = "unspecified"
  ssl                  = false
  ssl_verify           = true
}

resource "opnsense_haproxy_backend" "app" {
  enabled                  = false
  name                     = "example_app_backend"
  description              = "Example HAProxy backend pool"
  mode                     = "http"
  algorithm                = "source"
  random_draws             = "2"
  linked_servers           = [opnsense_haproxy_server.app.id]
  health_check_enabled     = false
  health_check_proxy_proto = "backend"
  http2_enabled            = true
  http2_enabled_nontls     = false
  ba_advertised_protocols  = ["h2", "http11"]
  persistence              = "sticktable"
  persistence_cookiemode   = "piggyback"
  persistence_cookiename   = "SRVCOOKIE"
  persistence_stripquotes  = true
  stickiness_pattern       = "sourceipv4"
  stickiness_expire        = "30m"
  stickiness_size          = "50k"
  tuning_noport            = false
  tuning_httpreuse         = "safe"
  tuning_caching           = false
}
