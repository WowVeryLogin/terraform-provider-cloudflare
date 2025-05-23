resource "cloudflare_workers_kv_namespace" "%[1]s" {
	account_id = "%[4]s"
	title = "%[1]s"
}

resource "cloudflare_queue" "%[1]s" {
	account_id = "%[4]s"
	queue_name = "%[1]s"
}

resource "cloudflare_workers_script" "%[1]s" {
  account_id = "%[4]s"
  script_name    = "%[1]s"
  content = "%[2]s"

  bindings = [
    {
      name = "MY_KV_NAMESPACE"
      type = "kv_namespace"
      namespace_id = cloudflare_workers_kv_namespace.%[1]s.id
    },
    {
      name         = "MY_QUEUE"
      type = "queue"
      queue_name = cloudflare_queue.%[1]s.queue_name
    }
  ]
}
