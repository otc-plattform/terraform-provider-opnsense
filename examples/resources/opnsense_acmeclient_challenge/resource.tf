resource "opnsense_acmeclient_challenge" "example_http" {
  enabled                    = true
  name                       = "tf-example-http"
  description                = "Terraform HTTP-01 challenge"
  method                     = "http01"
  dns_service                = "dns_freedns"
  http_service               = "opnsense"
  http_opn_autodiscovery     = true
  tlsalpn_acme_autodiscovery = true
}
