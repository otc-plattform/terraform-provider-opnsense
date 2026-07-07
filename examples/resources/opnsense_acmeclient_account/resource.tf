resource "opnsense_acmeclient_account" "example" {
  enabled     = true
  name        = "tf-example-account"
  email       = "admin@example.com"
  ca          = "letsencrypt"
  description = "Terraform test account"
}
