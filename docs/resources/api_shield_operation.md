---
page_title: "cloudflare_api_shield_operation Resource - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_api_shield_operation (Resource)



## Example Usage

```terraform
resource "cloudflare_api_shield_operation" "example" {
  zone_id  = "0da42c8d2132a9ddaf714f9e7c920711"
  method   = "GET"
  host     = "api.example.com"
  endpoint = "/path"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `operation_id` (String) UUID identifier
- `zone_id` (String) Identifier

### Optional

- `state` (String) Mark state of operation in API Discovery
  * `review` - Mark operation as for review
  * `ignored` - Mark operation as ignored

