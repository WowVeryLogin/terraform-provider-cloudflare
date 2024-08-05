// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package static_route

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r StaticRouteResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description:   "Identifier",
						Required:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"route_id": schema.StringAttribute{
						Description:   "Identifier",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"description": schema.StringAttribute{
						Description: "An optional human provided description of the static route.",
						Optional:    true,
					},
					"nexthop": schema.StringAttribute{
						Description: "The next-hop IP Address for the static route.",
						Optional:    true,
					},
					"prefix": schema.StringAttribute{
						Description: "IP Prefix in Classless Inter-Domain Routing format.",
						Optional:    true,
					},
					"priority": schema.Int64Attribute{
						Description: "Priority of the static route.",
						Optional:    true,
					},
					"weight": schema.Int64Attribute{
						Description: "Optional weight of the ECMP scope - if provided.",
						Optional:    true,
					},
					"scope": schema.SingleNestedAttribute{
						Description: "Used only for ECMP routes.",
						Optional:    true,
						Attributes: map[string]schema.Attribute{
							"colo_names": schema.ListAttribute{
								Description: "List of colo names for the ECMP scope.",
								Optional:    true,
								ElementType: types.StringType,
							},
							"colo_regions": schema.ListAttribute{
								Description: "List of colo regions for the ECMP scope.",
								Optional:    true,
								ElementType: types.StringType,
							},
						},
					},
					"deleted": schema.BoolAttribute{
						Computed: true,
					},
					"modified": schema.BoolAttribute{
						Computed: true,
					},
					"routes": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"nexthop": schema.StringAttribute{
									Description: "The next-hop IP Address for the static route.",
									Required:    true,
								},
								"prefix": schema.StringAttribute{
									Description: "IP Prefix in Classless Inter-Domain Routing format.",
									Required:    true,
								},
								"priority": schema.Int64Attribute{
									Description: "Priority of the static route.",
									Required:    true,
								},
								"id": schema.StringAttribute{
									Description: "Identifier",
									Optional:    true,
								},
								"created_on": schema.StringAttribute{
									Description: "When the route was created.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"description": schema.StringAttribute{
									Description: "An optional human provided description of the static route.",
									Optional:    true,
								},
								"modified_on": schema.StringAttribute{
									Description: "When the route was last modified.",
									Computed:    true,
									CustomType:  timetypes.RFC3339Type{},
								},
								"scope": schema.SingleNestedAttribute{
									Description: "Used only for ECMP routes.",
									Optional:    true,
									Attributes: map[string]schema.Attribute{
										"colo_names": schema.ListAttribute{
											Description: "List of colo names for the ECMP scope.",
											Optional:    true,
											ElementType: types.StringType,
										},
										"colo_regions": schema.ListAttribute{
											Description: "List of colo regions for the ECMP scope.",
											Optional:    true,
											ElementType: types.StringType,
										},
									},
								},
								"weight": schema.Int64Attribute{
									Description: "Optional weight of the ECMP scope - if provided.",
									Optional:    true,
								},
							},
						},
					},
					"deleted_route": schema.StringAttribute{
						Computed: true,
					},
					"modified_route": schema.StringAttribute{
						Computed: true,
					},
					"route": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state StaticRouteModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}