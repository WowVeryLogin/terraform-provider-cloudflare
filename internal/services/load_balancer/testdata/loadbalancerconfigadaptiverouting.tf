resource "cloudflare_load_balancer" "%[3]s" {
  zone_id = "%[1]s"
  name = "tf-testacc-lb-adaptive-routing-%[3]s.%[2]s"
  fallback_pool = "${cloudflare_load_balancer_pool.%[3]s.id}"
  default_pools = ["${cloudflare_load_balancer_pool.%[3]s.id}"]
  adaptive_routing = {
    failover_across_pools = true
  }
  session_affinity = "cookie"
}
