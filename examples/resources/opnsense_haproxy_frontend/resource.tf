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
  config = {
    enabled                 = "0"
    name                    = "example_frontend_backend"
    mode                    = "http"
    algorithm               = "source"
    random_draws            = "2"
    linkedServers           = opnsense_haproxy_server.app.id
    healthCheckEnabled      = "0"
    healthCheckProxyProto   = "backend"
    http2Enabled            = "1"
    ba_advertised_protocols = "h2,http11"
    persistence             = ""
    tuning_httpreuse        = "safe"
  }
}

resource "opnsense_haproxy_frontend" "app" {
  config = {
    enabled                   = "0"
    name                      = "example_public_service"
    description               = "Example HAProxy public service"
    bind                      = "0.0.0.0:8080"
    bindOptions               = ""
    mode                      = "http"
    defaultBackend            = opnsense_haproxy_backend.app.id
    ssl_enabled               = "0"
    ssl_advancedEnabled       = "0"
    basicAuthEnabled          = "0"
    logging_dontLogNull       = "0"
    logging_dontLogNormal     = "0"
    logging_logSeparateErrors = "0"
    logging_detailedLog       = "0"
    logging_socketStats       = "0"
    http2Enabled              = "1"
    http2Enabled_nontls       = "0"
    advertised_protocols      = "h2,http11"
    forwardFor                = "0"
    prometheus_enabled        = "0"
    prometheus_path           = "/metrics"
    connectionBehaviour       = "http-keep-alive"
    customOptions             = ""
    linkedActions             = ""
    linkedErrorfiles          = ""
  }
}
