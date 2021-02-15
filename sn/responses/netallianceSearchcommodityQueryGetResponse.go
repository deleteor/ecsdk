package responses

type NetallianceSearchcommodityQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			QuerySearchcommodity []QueryCommoditydetail `json:"querySearchcommodity"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}
