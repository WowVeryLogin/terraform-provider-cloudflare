// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_device_default_profile

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*ZeroTrustDeviceDefaultProfileDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_id": schema.StringAttribute{
				Required: true,
			},
			"allow_mode_switch": schema.BoolAttribute{
				Description: "Whether to allow the user to switch WARP between modes.",
				Computed:    true,
			},
			"allow_updates": schema.BoolAttribute{
				Description: "Whether to receive update notifications when a new version of the client is available.",
				Computed:    true,
			},
			"allowed_to_leave": schema.BoolAttribute{
				Description: "Whether to allow devices to leave the organization.",
				Computed:    true,
			},
			"auto_connect": schema.Float64Attribute{
				Description: "The amount of time in seconds to reconnect after having been disabled.",
				Computed:    true,
			},
			"captive_portal": schema.Float64Attribute{
				Description: "Turn on the captive portal after the specified amount of time.",
				Computed:    true,
			},
			"default": schema.BoolAttribute{
				Description: "Whether the policy will be applied to matching devices.",
				Computed:    true,
			},
			"disable_auto_fallback": schema.BoolAttribute{
				Description: "If the `dns_server` field of a fallback domain is not present, the client will fall back to a best guess of the default/system DNS resolvers unless this policy option is set to `true`.",
				Computed:    true,
			},
			"enabled": schema.BoolAttribute{
				Description: "Whether the policy will be applied to matching devices.",
				Computed:    true,
			},
			"exclude_office_ips": schema.BoolAttribute{
				Description: "Whether to add Microsoft IPs to Split Tunnel exclusions.",
				Computed:    true,
			},
			"gateway_unique_id": schema.StringAttribute{
				Computed: true,
			},
			"register_interface_ip_with_dns": schema.BoolAttribute{
				Description: "Determines if the operating system will register WARP's local interface IP with your on-premises DNS server.",
				Computed:    true,
			},
			"support_url": schema.StringAttribute{
				Description: "The URL to launch when the Send Feedback button is clicked.",
				Computed:    true,
			},
			"switch_locked": schema.BoolAttribute{
				Description: "Whether to allow the user to turn off the WARP switch and disconnect the client.",
				Computed:    true,
			},
			"tunnel_protocol": schema.StringAttribute{
				Description: "Determines which tunnel protocol to use.",
				Computed:    true,
			},
			"exclude": schema.ListNestedAttribute{
				Description: "List of routes excluded in the WARP client's tunnel.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceDefaultProfileExcludeDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "The address in CIDR format to exclude from the tunnel. If `address` is present, `host` must not be present.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the Split Tunnel item, displayed in the client UI.",
							Computed:    true,
						},
						"host": schema.StringAttribute{
							Description: "The domain name to exclude from the tunnel. If `host` is present, `address` must not be present.",
							Computed:    true,
						},
					},
				},
			},
			"fallback_domains": schema.ListNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectListType[ZeroTrustDeviceDefaultProfileFallbackDomainsDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"suffix": schema.StringAttribute{
							Description: "The domain suffix to match when resolving locally.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the fallback domain, displayed in the client UI.",
							Computed:    true,
						},
						"dns_server": schema.ListAttribute{
							Description: "A list of IP addresses to handle domain resolution.",
							Computed:    true,
							CustomType:  customfield.NewListType[types.String](ctx),
							ElementType: types.StringType,
						},
					},
				},
			},
			"include": schema.ListNestedAttribute{
				Description: "List of routes included in the WARP client's tunnel.",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[ZeroTrustDeviceDefaultProfileIncludeDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"address": schema.StringAttribute{
							Description: "The address in CIDR format to include in the tunnel. If `address` is present, `host` must not be present.",
							Computed:    true,
						},
						"description": schema.StringAttribute{
							Description: "A description of the Split Tunnel item, displayed in the client UI.",
							Computed:    true,
						},
						"host": schema.StringAttribute{
							Description: "The domain name to include in the tunnel. If `host` is present, `address` must not be present.",
							Computed:    true,
						},
					},
				},
			},
			"service_mode_v2": schema.SingleNestedAttribute{
				Computed:   true,
				CustomType: customfield.NewNestedObjectType[ZeroTrustDeviceDefaultProfileServiceModeV2DataSourceModel](ctx),
				Attributes: map[string]schema.Attribute{
					"mode": schema.StringAttribute{
						Description: "The mode to run the WARP client under.",
						Computed:    true,
					},
					"port": schema.Float64Attribute{
						Description: "The port number when used with proxy mode.",
						Computed:    true,
					},
				},
			},
		},
	}
}

func (d *ZeroTrustDeviceDefaultProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *ZeroTrustDeviceDefaultProfileDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
