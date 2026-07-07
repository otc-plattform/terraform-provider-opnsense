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
  config = {
    enabled                   = "0"
    name                      = "example_app_backend"
    description               = "Example HAProxy backend pool"
    mode                      = "http"
    algorithm                 = "source"
    random_draws              = "2"
    proxyProtocol             = ""
    linkedServers             = opnsense_haproxy_server.app.id
    linkedFcgi                = ""
    linkedResolver            = ""
    resolverOpts              = ""
    resolvePrefer             = ""
    source                    = ""
    healthCheckEnabled        = "0"
    healthCheck               = ""
    healthCheckLogStatus      = "0"
    healthCheckProxyProto     = "backend"
    http2Enabled              = "1"
    http2Enabled_nontls       = "0"
    ba_advertised_protocols   = "h2,http11"
    forwardFor                = "0"
    forwardedHeader           = "0"
    forwardedHeaderParameters = ""
    persistence               = ""
    customOptions             = ""
    tuning_noport             = "0"
    tuning_httpreuse          = "safe"
    tuning_caching            = "0"
    linkedActions             = ""
    linkedErrorfiles          = ""
  }
}
