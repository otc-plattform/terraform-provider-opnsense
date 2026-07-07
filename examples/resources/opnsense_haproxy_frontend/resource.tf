resource "opnsense_haproxy_server" "app" {
  enabled              = false
  name                 = "example_frontend_server"
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
  name                     = "example_frontend_backend"
  mode                     = "http"
  algorithm                = "source"
  random_draws             = "2"
  linked_servers           = [opnsense_haproxy_server.app.id]
  health_check_enabled     = false
  health_check_proxy_proto = "backend"
  http2_enabled            = true
  ba_advertised_protocols  = ["h2", "http11"]
  persistence              = "sticktable"
  tuning_httpreuse         = "safe"
}

resource "opnsense_haproxy_frontend" "app" {
  enabled                     = false
  name                        = "example_public_service"
  description                 = "Example HAProxy public service"
  bind                        = ["0.0.0.0:8080"]
  mode                        = "http"
  default_backend             = opnsense_haproxy_backend.app.id
  ssl_enabled                 = false
  ssl_advanced_enabled        = false
  basic_auth_enabled          = false
  logging_dont_log_null       = false
  logging_dont_log_normal     = false
  logging_log_separate_errors = false
  logging_detailed_log        = false
  logging_socket_stats        = false
  http2_enabled               = true
  http2_enabled_nontls        = false
  advertised_protocols        = ["h2", "http11"]
  forward_for                 = "0"
  prometheus_enabled          = false
  prometheus_path             = "/metrics"
  connection_behaviour        = "http-keep-alive"
}
