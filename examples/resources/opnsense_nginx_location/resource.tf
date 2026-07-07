resource "opnsense_nginx_location" "example" {
  description        = "tf-example-location"
  url_pattern        = "/app"
  match_type         = "="
  advanced_acl       = false
  cache_min_uses     = 1
  upstream_keepalive = false
}
