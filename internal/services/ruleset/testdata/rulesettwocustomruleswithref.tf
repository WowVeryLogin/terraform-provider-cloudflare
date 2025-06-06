
resource "cloudflare_ruleset" "%[1]s" {
  zone_id     = "%[2]s"
  name        = "Terraform provider test"
  description = "%[1]s ruleset description"
  kind        = "zone"
  phase       = "http_request_firewall_custom"
  rules = [{
    action     = "log"
    enabled    = %[3]t
    expression = "(http.request.uri.path eq \"/admin\")"
    ref        = "foo"
    },
    {
      action     = "challenge"
      enabled    = true
      expression = "(http.request.uri.path eq \"/login\")"
  }]
}