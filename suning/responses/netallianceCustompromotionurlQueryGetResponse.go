package responses

type NetallianceCustompromotionurlQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			QueryCustompromotionurl QueryCustompromotionurl`json:"queryCustompromotionurl"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}

type QueryCustompromotionurl struct {
	ExtendUrl  string  `json:"extendUrl"`  //推广链接(获取到链接后，需进行URL解码才可推广；如果需要加入子会员id，解码后的链接中增加字段sub_user)
	ShortUrl string `json:"shortUrl"` //推广短链接(获取到链接后，需进行URL解码才可推广；如果需要加入子会员id，解码后的链接中增加字段sub_user)
}