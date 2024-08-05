---
page_title: "cloudflare_waiting_room_event Data Source - Cloudflare"
subcategory: ""
description: |-
  
---

# cloudflare_waiting_room_event (Data Source)




<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `custom_page_html` (String) If set, the event will override the waiting room's `custom_page_html` property while it is active. If null, the event will inherit it.
- `disable_session_renewal` (Boolean) If set, the event will override the waiting room's `disable_session_renewal` property while it is active. If null, the event will inherit it.
- `event_end_time` (String) An ISO 8601 timestamp that marks the end of the event.
- `event_id` (String)
- `event_start_time` (String) An ISO 8601 timestamp that marks the start of the event. At this time, queued users will be processed with the event's configuration. The start time must be at least one minute before `event_end_time`.
- `filter` (Attributes) (see [below for nested schema](#nestedatt--filter))
- `name` (String) A unique name to identify the event. Only alphanumeric characters, hyphens and underscores are allowed.
- `new_users_per_minute` (Number) If set, the event will override the waiting room's `new_users_per_minute` property while it is active. If null, the event will inherit it. This can only be set if the event's `total_active_users` property is also set.
- `prequeue_start_time` (String) An ISO 8601 timestamp that marks when to begin queueing all users before the event starts. The prequeue must start at least five minutes before `event_start_time`.
- `queueing_method` (String) If set, the event will override the waiting room's `queueing_method` property while it is active. If null, the event will inherit it.
- `session_duration` (Number) If set, the event will override the waiting room's `session_duration` property while it is active. If null, the event will inherit it.
- `total_active_users` (Number) If set, the event will override the waiting room's `total_active_users` property while it is active. If null, the event will inherit it. This can only be set if the event's `new_users_per_minute` property is also set.
- `waiting_room_id` (String)
- `zone_id` (String) Identifier

### Read-Only

- `created_on` (String)
- `description` (String) A note that you can use to add more details about the event.
- `id` (String) The ID of this resource.
- `modified_on` (String)
- `shuffle_at_event_start` (Boolean) If enabled, users in the prequeue will be shuffled randomly at the `event_start_time`. Requires that `prequeue_start_time` is not null. This is useful for situations when many users will join the event prequeue at the same time and you want to shuffle them to ensure fairness. Naturally, it makes the most sense to enable this feature when the `queueing_method` during the event respects ordering such as **fifo**, or else the shuffling may be unnecessary.
- `suspended` (Boolean) Suspends or allows an event. If set to `true`, the event is ignored and traffic will be handled based on the waiting room configuration.

<a id="nestedatt--filter"></a>
### Nested Schema for `filter`

Required:

- `waiting_room_id` (String)
- `zone_id` (String) Identifier

Optional:

- `page` (String) Page number of paginated results.
- `per_page` (String) Maximum number of results per page. Must be a multiple of 5.

