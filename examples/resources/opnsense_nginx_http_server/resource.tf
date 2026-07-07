resource "opnsense_nginx_location" "app" {
  description = "tf-example-location"
  url_pattern = "/app"
  match_type  = "="
}

resource "opnsense_nginx_http_server" "example" {
  server_name          = "tf-http-server"
  listen_http_address  = "0.0.0.0:8080"
  default_server       = true
  https_only           = false
  enable_acme_support  = false
  log_handshakes       = false
  tls_reject_handshake = false
  locations            = [opnsense_nginx_location.app.id]
  tls_protocols        = ["TLSv1.2", "TLSv1.3"]
}
