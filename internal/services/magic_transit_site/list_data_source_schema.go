// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit_site

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*MagicTransitSitesDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Description: "Identifier",
				Required:    true,
			},
			"connector_identifier": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[MagicTransitSitesResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"connector_id": schema.StringAttribute{
							Description: "Magic Connector identifier tag.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"ha_mode": schema.BoolAttribute{
							Description: "Site high availability mode. If set to true, the site can have two connectors and runs in high availability mode.",
							Computed:    true,
						},
						"location": schema.SingleNestedAttribute{
							Description: "Location of site in latitude and longitude.",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[MagicTransitSitesLocationDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"lat": schema.StringAttribute{
									Description: "Latitude",
									Computed:    true,
								},
								"lon": schema.StringAttribute{
									Description: "Longitude",
									Computed:    true,
								},
							},
						},
						"name": schema.StringAttribute{
							Description: "The name of the site.",
							Computed:    true,
						},
						"secondary_connector_id": schema.StringAttribute{
							Description: "Magic Connector identifier tag. Used when high availability mode is on.",
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func (d *MagicTransitSitesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *MagicTransitSitesDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}