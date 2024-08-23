// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust_dlp_custom_profile

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/zero_trust"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type ZeroTrustDLPCustomProfileResultDataSourceEnvelope struct {
	Result ZeroTrustDLPCustomProfileDataSourceModel `json:"result,computed"`
}

type ZeroTrustDLPCustomProfileDataSourceModel struct {
	AccountID types.String `tfsdk:"account_id" path:"account_id"`
	ProfileID types.String `tfsdk:"profile_id" path:"profile_id"`
}

func (m *ZeroTrustDLPCustomProfileDataSourceModel) toReadParams() (params zero_trust.DLPProfileCustomGetParams, diags diag.Diagnostics) {
	params = zero_trust.DLPProfileCustomGetParams{
		AccountID: cloudflare.F(m.AccountID.ValueString()),
	}

	return
}