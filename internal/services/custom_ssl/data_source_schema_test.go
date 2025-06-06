// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package custom_ssl_test

import (
	"context"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/services/custom_ssl"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/test_helpers"
)

func TestCustomSSLDataSourceModelSchemaParity(t *testing.T) {
	t.Parallel()
	model := (*custom_ssl.CustomSSLDataSourceModel)(nil)
	schema := custom_ssl.DataSourceSchema(context.TODO())
	errs := test_helpers.ValidateDataSourceModelSchemaIntegrity(model, schema)
	errs.Report(t)
}
