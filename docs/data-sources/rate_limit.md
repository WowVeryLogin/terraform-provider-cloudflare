---
page_title: "cloudflare_rate_limit Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_rate_limit (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `action` (Attributes) The action to perform when the threshold of matched traffic within the configured period is exceeded. (see [below for nested schema](#nestedatt--action))
- `bypass` (Attributes List) Criteria specifying when the current rate limit should be bypassed. You can specify that the rate limit should not apply to one or more URLs. (see [below for nested schema](#nestedatt--bypass))
- `description` (String) An informative summary of the rate limit. This value is sanitized and any tags will be removed.
- `disabled` (Boolean) When true, indicates that the rate limit is currently disabled.
- `find_one_by` (Attributes) (see [below for nested schema](#nestedatt--find_one_by))
- `id` (String) The unique identifier of the rate limit.
- `match` (Attributes) Determines which traffic the rate limit counts towards the threshold. (see [below for nested schema](#nestedatt--match))
- `period` (Number) The time in seconds (an integer value) to count matching traffic. If the count exceeds the configured threshold within this period, Cloudflare will perform the configured action.
- `threshold` (Number) The threshold that will trigger the configured mitigation action. Configure this value along with the `period` property to establish a threshold per period.
- `zone_identifier` (String) Identifier

<a id="nestedatt--action"></a>
### Nested Schema for `action`

Optional:

- `mode` (String) The action to perform.
- `response` (Attributes) A custom content type and reponse to return when the threshold is exceeded. The custom response configured in this object will override the custom error for the zone. This object is optional.
Notes: If you omit this object, Cloudflare will use the default HTML error page. If "mode" is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the zone challenge pages and you should not provide the "response" object. (see [below for nested schema](#nestedatt--action--response))
- `timeout` (Number) The time in seconds during which Cloudflare will perform the mitigation action. Must be an integer value greater than or equal to the period.
Notes: If "mode" is "challenge", "managed_challenge", or "js_challenge", Cloudflare will use the zone's Challenge Passage time and you should not provide this value.

<a id="nestedatt--action--response"></a>
### Nested Schema for `action.response`

Optional:

- `body` (String) The response body to return. The value must conform to the configured content type.
- `content_type` (String) The content type of the body. Must be one of the following: `text/plain`, `text/xml`, or `application/json`.



<a id="nestedatt--bypass"></a>
### Nested Schema for `bypass`

Optional:

- `name` (String)
- `value` (String) The URL to bypass.


<a id="nestedatt--find_one_by"></a>
### Nested Schema for `find_one_by`

Required:

- `zone_identifier` (String) Identifier

Optional:

- `page` (Number) The page number of paginated results.
- `per_page` (Number) The maximum number of results per page. You can only set the value to `1` or to a multiple of 5 such as `5`, `10`, `15`, or `20`.


<a id="nestedatt--match"></a>
### Nested Schema for `match`

Optional:

- `headers` (Attributes List) (see [below for nested schema](#nestedatt--match--headers))
- `request` (Attributes) (see [below for nested schema](#nestedatt--match--request))
- `response` (Attributes) (see [below for nested schema](#nestedatt--match--response))

<a id="nestedatt--match--headers"></a>
### Nested Schema for `match.headers`

Optional:

- `name` (String) The name of the response header to match.
- `op` (String) The operator used when matching: `eq` means "equal" and `ne` means "not equal".
- `value` (String) The value of the response header, which must match exactly.


<a id="nestedatt--match--request"></a>
### Nested Schema for `match.request`

Optional:

- `methods` (List of String) The HTTP methods to match. You can specify a subset (for example, `['POST','PUT']`) or all methods (`['_ALL_']`). This field is optional when creating a rate limit.
- `schemes` (List of String) The HTTP schemes to match. You can specify one scheme (`['HTTPS']`), both schemes (`['HTTP','HTTPS']`), or all schemes (`['_ALL_']`). This field is optional.
- `url` (String) The URL pattern to match, composed of a host and a path such as `example.org/path*`. Normalization is applied before the pattern is matched. `*` wildcards are expanded to match applicable traffic. Query strings are not matched. Set the value to `*` to match all traffic to your zone.


<a id="nestedatt--match--response"></a>
### Nested Schema for `match.response`

Optional:

- `origin_traffic` (Boolean) When true, only the uncached traffic served from your origin servers will count towards rate limiting. In this case, any cached traffic served by Cloudflare will not count towards rate limiting. This field is optional.
Notes: This field is deprecated. Instead, use response headers and set "origin_traffic" to "false" to avoid legacy behaviour interacting with the "response_headers" property.

