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
  config = {
    enabled                 = "0"
    name                    = "example_acme_backend"
    mode                    = "http"
    algorithm               = "source"
    random_draws            = "2"
    linkedServers           = opnsense_haproxy_server.acme.id
    healthCheckEnabled      = "0"
    healthCheckProxyProto   = "backend"
    http2Enabled            = "1"
    ba_advertised_protocols = "h2,http11"
    persistence             = ""
    tuning_httpreuse        = "safe"
  }
}

resource "opnsense_haproxy_acl" "acme_challenge" {
  config = {
    name          = "example_acme_challenge"
    description   = "Match ACME HTTP-01 challenge requests"
    expression    = "path_beg"
    negate        = "0"
    caseSensitive = "0"
    path_beg      = "/.well-known/acme-challenge/"
  }
}

resource "opnsense_haproxy_action" "acme_challenge" {
  config = {
    enabled     = "0"
    name        = "example_acme_redirect"
    description = "Route ACME challenge requests to the ACME backend"
    testType    = "if"
    linkedAcls  = opnsense_haproxy_acl.acme_challenge.id
    operator    = "and"
    type        = "use_backend"
    use_backend = opnsense_haproxy_backend.acme.id
    use_server  = ""
  }
}
