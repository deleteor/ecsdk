package responses

type NetallianceExtensionlinkGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			GetExtensionlink GetExtensionlink `json:"getExtensionlink"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}

type GetExtensionlink struct {
	ShortLink string `json:"shortLink"` //生成推广的短链接
}