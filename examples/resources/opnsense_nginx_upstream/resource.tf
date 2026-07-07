resource "opnsense_nginx_upstream_server" "app" {
  description = "tf-upstream-server"
  server      = "10.0.0.10"
  port        = "8080"
}

resource "opnsense_nginx_upstream" "example" {
  description              = "tf-upstream"
  load_balancing_algorithm = ""
  proxy_protocol           = false
  tls_enable               = false
  server_entries           = [opnsense_nginx_upstream_server.app.id]
}
