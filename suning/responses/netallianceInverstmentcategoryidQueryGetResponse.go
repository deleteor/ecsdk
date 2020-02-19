package responses

type NetallianceInverstmentcategoryidQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			ResponseCategory struct {
				CategoryName string         `json:"categoryName"`
				CategoryList []CategoryList `json:"categoryList"`
			} `json:"queryInverstmentcategoryid"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}

type CategoryList struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
