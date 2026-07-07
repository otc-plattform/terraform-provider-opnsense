resource "opnsense_nginx_upstream_server" "example" {
  description = "tf-upstream-server"
  server      = "10.0.0.10"
  port        = "8080"
  priority    = "1"
}
