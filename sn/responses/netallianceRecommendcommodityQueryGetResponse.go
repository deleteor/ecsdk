package responses

type NetallianceRecommendcommodityQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			QueryRecommendcommodity []QueryCommoditydetail `json:"queryRecommendcommodity"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}
