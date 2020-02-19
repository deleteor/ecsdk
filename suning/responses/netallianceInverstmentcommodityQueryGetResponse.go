package responses

type NetallianceInverstmentcommodityQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			QueryInverstmentcommodity []QueryCommoditydetail `json:"queryInverstmentcommodity"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}
