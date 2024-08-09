// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package logpull_retention

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type LogpullRetentionResultEnvelope struct {
	Result LogpullRetentionModel `json:"result,computed"`
}

type LogpullRetentionModel struct {
	ZoneID types.String `tfsdk:"zone_id" path:"zone_id"`
	Flag   types.Bool   `tfsdk:"flag" json:"flag"`
}