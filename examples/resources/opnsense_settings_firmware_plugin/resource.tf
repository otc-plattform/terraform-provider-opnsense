resource "opnsense_settings_firmware_plugin" "os_nginx" {
  package   = "os-nginx"
  installed = true
}
