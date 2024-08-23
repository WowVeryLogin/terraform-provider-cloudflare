// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package waiting_room_setting

import (
	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/waiting_rooms"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type WaitingRoomSettingResultDataSourceEnvelope struct {
	Result WaitingRoomSettingDataSourceModel `json:"result,computed"`
}

type WaitingRoomSettingDataSourceModel struct {
	ZoneID                    types.String `tfsdk:"zone_id" path:"zone_id"`
	SearchEngineCrawlerBypass types.Bool   `tfsdk:"search_engine_crawler_bypass" json:"search_engine_crawler_bypass"`
}

func (m *WaitingRoomSettingDataSourceModel) toReadParams() (params waiting_rooms.SettingGetParams, diags diag.Diagnostics) {
	params = waiting_rooms.SettingGetParams{
		ZoneID: cloudflare.F(m.ZoneID.ValueString()),
	}

	return
}