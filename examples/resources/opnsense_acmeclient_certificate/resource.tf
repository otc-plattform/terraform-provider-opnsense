resource "opnsense_acmeclient_account" "example" {
  enabled     = true
  name        = "tf-example-account"
  email       = "admin@example.com"
  ca          = "letsencrypt"
  description = "Terraform test account"
}

resource "opnsense_acmeclient_challenge" "http" {
  enabled                    = true
  name                       = "tf-example-http"
  method                     = "http01"
  http_service               = "opnsense"
  http_opn_autodiscovery     = true
  tlsalpn_acme_autodiscovery = true
}

resource "opnsense_acmeclient_automation" "restart_gui" {
  enabled = true
  name    = "tf-example-automation"
  type    = "configd_restart_gui"
}

resource "opnsense_acmeclient_certificate" "example" {
  enabled              = false
  name                 = "tf-example-cert"
  description          = "Terraform managed certificate"
  account_id           = opnsense_acmeclient_account.example.id
  validation_method_id = opnsense_acmeclient_challenge.http.id
  alt_names            = ["example.com", "www.example.com"]
  restart_actions      = [opnsense_acmeclient_automation.restart_gui.id]
  auto_renewal         = true
  renew_interval       = 30
  key_length           = "key_2048"
}
