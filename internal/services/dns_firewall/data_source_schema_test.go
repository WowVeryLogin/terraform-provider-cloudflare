// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns_firewall_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/dns_firewall"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestDNSFirewallDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*dns_firewall.DNSFirewallDataSourceModel)(nil)
	schema := dns_firewall.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
