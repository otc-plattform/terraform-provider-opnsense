resource "opnsense_settings_cron" "nightly_run" {
  enabled     = false
  minutes     = "37"
  hours       = "2"
  days        = "*"
  months      = "*"
  weekdays    = "*"
  who         = "root"
  command     = "firmware auto-update"
  parameters  = ""
  description = "Run nightly system backup at 2:30 AM"
}
