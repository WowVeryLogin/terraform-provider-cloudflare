// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package teams_proxy_endpoint

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TeamsProxyEndpointResultDataSourceEnvelope struct {
	Result TeamsProxyEndpointDataSourceModel `json:"result,computed"`
}

type TeamsProxyEndpointDataSourceModel struct {
	AccountID       types.String `tfsdk:"account_id" path:"account_id"`
	ProxyEndpointID types.String `tfsdk:"proxy_endpoint_id" path:"proxy_endpoint_id"`
}