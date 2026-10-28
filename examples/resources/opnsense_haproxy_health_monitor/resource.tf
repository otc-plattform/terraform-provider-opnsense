resource "opnsense_haproxy_health_monitor" "http" {
  name                    = "http-health"
  type                    = "http"
  interval                = "5s"
  http_method             = "get"
  http_uri                = "/health"
  http_expression_enabled = true
  http_expression         = "status"
  http_value              = "200"
}

# Add the monitor to the existing HAProxy apply resource along with every
# other HAProxy resource that must be validated and activated together.
resource "opnsense_haproxy_apply" "configuration" {
  triggers = {
    health_monitor = sha256(jsonencode(opnsense_haproxy_health_monitor.http))
  }

  depends_on = [opnsense_haproxy_health_monitor.http]
}
