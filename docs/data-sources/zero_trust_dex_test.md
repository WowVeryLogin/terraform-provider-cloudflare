---
page_title: "cloudflare_zero_trust_dex_test Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_zero_trust_dex_test (Data Source)



## Example Usage

```terraform
data "cloudflare_zero_trust_dex_test" "example_zero_trust_dex_test" {
  account_id = "699d98642c564d2e855e9661899b7252"
  dex_test_id = "372e67954025e0ba6aaa6d586b9e0b59"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `account_id` (String)

### Optional

- `dex_test_id` (String) The unique identifier for the test.

### Read-Only

- `data` (Attributes) The configuration object which contains the details for the WARP client to conduct the test. (see [below for nested schema](#nestedatt--data))
- `description` (String) Additional details about the test.
- `enabled` (Boolean) Determines whether or not the test is active.
- `id` (String) The unique identifier for the test.
- `interval` (String) How often the test will run.
- `name` (String) The name of the DEX test. Must be unique.
- `target_policies` (Attributes List) Device settings profiles targeted by this test (see [below for nested schema](#nestedatt--target_policies))
- `targeted` (Boolean)
- `test_id` (String) The unique identifier for the test.

<a id="nestedatt--data"></a>
### Nested Schema for `data`

Read-Only:

- `host` (String) The desired endpoint to test.
- `kind` (String) The type of test.
- `method` (String) The HTTP request method type.


<a id="nestedatt--target_policies"></a>
### Nested Schema for `target_policies`

Read-Only:

- `default` (Boolean) Whether the profile is the account default
- `id` (String) The id of the device settings profile
- `name` (String) The name of the device settings profile


