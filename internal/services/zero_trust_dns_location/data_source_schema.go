// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dns_location

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDNSLocationDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Optional: true,
			},
			"location_id": schema.StringAttribute{
				Optional: true,
			},
			"created_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"updated_at": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"client_default": schema.BoolAttribute{
				Description: "True if the location is the default location.",
				Computed:    true,
				Optional:    true,
			},
			"dns_destination_ips_id": schema.StringAttribute{
				Description: "The identifier of the pair of IPv4 addresses assigned to this location.",
				Computed:    true,
				Optional:    true,
			},
			"doh_subdomain": schema.StringAttribute{
				Description: "The DNS over HTTPS domain to send DNS requests to. This field is auto-generated by Gateway.",
				Computed:    true,
				Optional:    true,
			},
			"ecs_support": schema.BoolAttribute{
				Description: "True if the location needs to resolve EDNS queries.",
				Computed:    true,
				Optional:    true,
			},
			"id": schema.StringAttribute{
				Computed: true,
				Optional: true,
			},
			"ip": schema.StringAttribute{
				Description: "IPV6 destination ip assigned to this location. DNS requests sent to this IP will counted as the request under this location. This field is auto-generated by Gateway.",
				Computed:    true,
				Optional:    true,
			},
			"ipv4_destination": schema.StringAttribute{
				Description: "The primary destination IPv4 address from the pair identified by the dns_destination_ips_id. This field is read-only.",
				Computed:    true,
				Optional:    true,
			},
			"ipv4_destination_backup": schema.StringAttribute{
				Description: "The backup destination IPv4 address from the pair identified by the dns_destination_ips_id. This field is read-only.",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "The name of the location.",
				Computed:    true,
				Optional:    true,
			},
			"endpoints": schema.SingleNestedAttribute{
				Description: "The destination endpoints configured for this location. When updating a location, if this field is absent or set with null, the endpoints configuration remains unchanged.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"doh": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "A list of allowed source IP network ranges for this endpoint. When empty, all source IPs are allowed. A non-empty list is only effective if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "The IP address or IP CIDR.",
											Computed:    true,
										},
									},
								},
							},
							"require_token": schema.BoolAttribute{
								Description: "True if the endpoint requires [user identity](https://developers.cloudflare.com/cloudflare-one/connections/connect-devices/agentless/dns/dns-over-https/#filter-doh-requests-by-user) authentication.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"dot": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "A list of allowed source IP network ranges for this endpoint. When empty, all source IPs are allowed. A non-empty list is only effective if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "The IP address or IP CIDR.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
					"ipv4": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
						},
					},
					"ipv6": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						Attributes: map[string]schema.Attribute{
							"enabled": schema.BoolAttribute{
								Description: "True if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
							},
							"networks": schema.ListNestedAttribute{
								Description: "A list of allowed source IPv6 network ranges for this endpoint. When empty, all source IPs are allowed. A non-empty list is only effective if the endpoint is enabled for this location.",
								Computed:    true,
								Optional:    true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"network": schema.StringAttribute{
											Description: "The IPv6 address or IPv6 CIDR.",
											Computed:    true,
										},
									},
								},
							},
						},
					},
				},
			},
			"networks": schema.ListNestedAttribute{
				Description: "A list of network ranges that requests from this location would originate from. A non-empty list is only effective if the ipv4 endpoint is enabled for this location.",
				Computed:    true,
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"network": schema.StringAttribute{
							Description: "The IPv4 address or IPv4 CIDR. IPv4 CIDRs are limited to a maximum of /24.",
							Computed:    true,
						},
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Required: true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDNSLocationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDNSLocationDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("account_id"), path.MatchRoot("location_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("account_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("location_id")),
	}
}
