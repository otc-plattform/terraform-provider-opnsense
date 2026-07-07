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
