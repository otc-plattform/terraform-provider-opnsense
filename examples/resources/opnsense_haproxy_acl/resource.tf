resource "opnsense_haproxy_acl" "acme_challenge" {
  name           = "example_acme_challenge"
  description    = "Match ACME HTTP-01 challenge requests"
  expression     = "path_beg"
  negate         = false
  case_sensitive = false
  path_beg       = "/.well-known/acme-challenge/"
}
