// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package access_rule

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func (r AccessRuleResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	return map[int64]resource.StateUpgrader{
		0: {
			PriorSchema: &schema.Schema{
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						Description:   "The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"identifier": schema.StringAttribute{
						Description:   "The unique identifier of the resource.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"zone_id": schema.StringAttribute{
						Description:   "The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.",
						Optional:      true,
						PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
					},
					"mode": schema.StringAttribute{
						Description: "The action to apply to a matched request.",
						Required:    true,
						Validators: []validator.String{
							stringvalidator.OneOfCaseInsensitive("block", "challenge", "whitelist", "js_challenge", "managed_challenge"),
						},
					},
					"configuration": schema.SingleNestedAttribute{
						Description: "The rule configuration.",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"target": schema.StringAttribute{
								Description: "The configuration target. You must set the target to `ip` when specifying an IP address in the rule.",
								Optional:    true,
								Validators: []validator.String{
									stringvalidator.OneOfCaseInsensitive("ip", "ip6", "ip_range", "asn", "country"),
								},
							},
							"value": schema.StringAttribute{
								Description: "The IP address to match. This address will be compared to the IP address of incoming requests.",
								Optional:    true,
							},
						},
					},
					"notes": schema.StringAttribute{
						Description: "An informative summary of the rule, typically used as a reminder or explanation.",
						Optional:    true,
					},
					"id": schema.StringAttribute{
						Description: "Identifier",
						Computed:    true,
					},
				},
			},

			StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
				var state AccessRuleModel

				resp.Diagnostics.Append(req.State.Get(ctx, &state)...)

				if resp.Diagnostics.HasError() {
					return
				}

				resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
			},
		},
	}
}