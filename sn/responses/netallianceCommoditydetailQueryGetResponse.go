package responses

type NetallianceCommoditydetailQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			QueryCommoditydetail []QueryCommoditydetail `json:"queryCommoditydetail"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}
