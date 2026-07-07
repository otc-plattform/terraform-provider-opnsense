resource "opnsense_haproxy_settings" "example" {
  enabled                = false
  graceful_stop          = false
  hard_stop_after        = "60s"
  seamless_reload        = true
  show_intro             = true
  tuning_nbthread        = 1
  tuning_ssl_min_version = "TLSv1.2"
  tuning_ssl_bind_options = [
    "prefer-client-ciphers",
  ]

  defaults_timeout_client  = "30s"
  defaults_timeout_connect = "30s"
  defaults_timeout_server  = "30s"
  defaults_retries         = 3
  defaults_redispatch      = "x-1"
  defaults_init_addr = [
    "last",
    "libc",
  ]

  logging_host     = "127.0.0.1"
  logging_facility = "local0"
  logging_level    = "info"

  stats_enabled         = false
  stats_port            = 8822
  stats_prometheus_bind = "*:8404"
  stats_prometheus_path = "/metrics"

  cache_enabled               = false
  cache_total_max_size        = 4
  cache_max_age               = 60
  cache_max_secondary_entries = 10

  peers_enabled = false
  peers_port1   = 1024
  peers_port2   = 1024
}
