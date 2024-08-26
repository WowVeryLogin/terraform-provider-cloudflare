// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package healthcheck

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/datasourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSourceWithConfigValidators = (*HealthcheckDataSource)(nil)

func DataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"healthcheck_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"zone_id": schema.StringAttribute{
				Description: "Identifier",
				Optional:    true,
			},
			"consecutive_fails": schema.Int64Attribute{
				Description: "The number of consecutive fails required from a health check before changing the health to unhealthy.",
				Computed:    true,
			},
			"consecutive_successes": schema.Int64Attribute{
				Description: "The number of consecutive successes required from a health check before changing the health to healthy.",
				Computed:    true,
			},
			"created_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"failure_reason": schema.StringAttribute{
				Description: "The current failure reason if status is unhealthy.",
				Computed:    true,
			},
			"id": schema.StringAttribute{
				Description: "Identifier",
				Computed:    true,
			},
			"interval": schema.Int64Attribute{
				Description: "The interval between each health check. Shorter intervals may give quicker notifications if the origin status changes, but will increase load on the origin as we check from multiple locations.",
				Computed:    true,
			},
			"modified_on": schema.StringAttribute{
				Computed:   true,
				CustomType: timetypes.RFC3339Type{},
			},
			"retries": schema.Int64Attribute{
				Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately.",
				Computed:    true,
			},
			"status": schema.StringAttribute{
				Description: "The current status of the origin server according to the health check.",
				Computed:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive(
						"unknown",
						"healthy",
						"unhealthy",
						"suspended",
					),
				},
			},
			"suspended": schema.BoolAttribute{
				Description: "If suspended, no health checks are sent to the origin.",
				Computed:    true,
			},
			"timeout": schema.Int64Attribute{
				Description: "The timeout (in seconds) before marking the health check as failed.",
				Computed:    true,
			},
			"type": schema.StringAttribute{
				Description: "The protocol to use for the health check. Currently supported protocols are 'HTTP', 'HTTPS' and 'TCP'.",
				Computed:    true,
			},
			"address": schema.StringAttribute{
				Description: "The hostname or IP address of the origin server to run health checks on.",
				Computed:    true,
				Optional:    true,
			},
			"description": schema.StringAttribute{
				Description: "A human-readable description of the health check.",
				Computed:    true,
				Optional:    true,
			},
			"name": schema.StringAttribute{
				Description: "A short name to identify the health check. Only alphanumeric characters, hyphens and underscores are allowed.",
				Computed:    true,
				Optional:    true,
			},
			"check_regions": schema.ListAttribute{
				Description: "A list of regions from which to run health checks. Null means Cloudflare will pick a default region.",
				Computed:    true,
				Optional:    true,
				Validators: []validator.List{
					listvalidator.ValueStringsAre(
						stringvalidator.OneOfCaseInsensitive(
							"WNAM",
							"ENAM",
							"WEU",
							"EEU",
							"NSAM",
							"SSAM",
							"OC",
							"ME",
							"NAF",
							"SAF",
							"IN",
							"SEAS",
							"NEAS",
							"ALL_REGIONS",
						),
					),
				},
				ElementType: types.StringType,
			},
			"http_config": schema.SingleNestedAttribute{
				Description: "Parameters specific to an HTTP or HTTPS health check.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"allow_insecure": schema.BoolAttribute{
						Description: "Do not validate the certificate when the health check uses HTTPS.",
						Computed:    true,
					},
					"expected_body": schema.StringAttribute{
						Description: "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy.",
						Computed:    true,
						Optional:    true,
					},
					"expected_codes": schema.ListAttribute{
						Description: "The expected HTTP response codes (e.g. \"200\") or code ranges (e.g. \"2xx\" for all codes starting with 2) of the health check.",
						Computed:    true,
						Optional:    true,
						ElementType: types.StringType,
					},
					"follow_redirects": schema.BoolAttribute{
						Description: "Follow redirects if the origin returns a 3xx status code.",
						Computed:    true,
					},
					"header": schema.MapAttribute{
						Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden.",
						Computed:    true,
						Optional:    true,
						ElementType: types.ListType{
							ElemType: types.StringType,
						},
					},
					"method": schema.StringAttribute{
						Description: "The HTTP method to use for the health check.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("GET", "HEAD"),
						},
					},
					"path": schema.StringAttribute{
						Description: "The endpoint path to health check against.",
						Computed:    true,
					},
					"port": schema.Int64Attribute{
						Description: "Port number to connect to for the health check. Defaults to 80 if type is HTTP or 443 if type is HTTPS.",
						Computed:    true,
					},
				},
			},
			"tcp_config": schema.SingleNestedAttribute{
				Description: "Parameters specific to TCP health check.",
				Computed:    true,
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"method": schema.StringAttribute{
						Description: "The TCP connection method to use for the health check.",
						Computed:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("connection_established"),
						},
					},
					"port": schema.Int64Attribute{
						Description: "Port number to connect to for the health check. Defaults to 80.",
						Computed:    true,
					},
				},
			},
			"filter": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"zone_id": schema.StringAttribute{
						Description: "Identifier",
						Required:    true,
					},
				},
			},
		},
	}
}

func (d *HealthcheckDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = DataSourceSchema(ctx)
}

func (d *HealthcheckDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{
		datasourcevalidator.RequiredTogether(path.MatchRoot("healthcheck_id"), path.MatchRoot("zone_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("healthcheck_id")),
		datasourcevalidator.ExactlyOneOf(path.MatchRoot("filter"), path.MatchRoot("zone_id")),
	}
}
