// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package page_rule

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type PageRulesDataSource struct {
	client *cloudflare.Client
}

var _ datasource.DataSourceWithConfigure = &PageRulesDataSource{}

func NewPageRulesDataSource() datasource.DataSource {
	return &PageRulesDataSource{}
}

func (d *PageRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_page_rules"
}

func (r *PageRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*cloudflare.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"unexpected resource configure type",
			fmt.Sprintf("Expected *cloudflare.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *PageRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *PageRulesDataSource

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}