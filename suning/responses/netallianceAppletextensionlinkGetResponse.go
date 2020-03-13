package responses

type NetallianceAppletextensionlinkGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			GetAppletextensionlink GetAppletextensionlink`json:"getAppletextensionlink"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}

type GetAppletextensionlink struct {
	AppId  string  `json:"appId"`  //小程序id
	AppletExtendUrl string `json:"appletExtendUrl"` //小程序推广链接pages/fourth/fourth/fourth+ "?actId=" + pgActionId + "&union=" + union;"
}