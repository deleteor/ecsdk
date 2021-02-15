package requests

type NetallianceAppletextensionlinkGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceAppletextensionlinkGetRequest() *NetallianceAppletextensionlinkGetRequest {
	req := &NetallianceAppletextensionlinkGetRequest{}
	req.Form = make(TopValues)
	return req
}

// SetProductUrl 商品的URL仅支持拼购和易购单品
func (req *NetallianceAppletextensionlinkGetRequest) SetProductUrl(productUrl string) {
	req.Form.Add("productUrl", productUrl)
}

// SetQuanUrl 券URL有时生成券推广，没有时生成商品推广
func (req *NetallianceAppletextensionlinkGetRequest) SetQuanUrl(quanUrl string) {
	req.Form.Add("quanUrl", quanUrl)
}

// SetPromotionId 这个推广只联盟前台申请的推广位，如果没有可以不填
func (req *NetallianceAppletextensionlinkGetRequest) SetPromotionId(promotionId string) {
	req.Form.Add("promotionId", promotionId)
}

// SetSubUser 子会员编码
func (req *NetallianceAppletextensionlinkGetRequest) SetSubUser(subUser string) {
	req.Form.Add("subUser", subUser)
}

// APIName 获取接口的名称
func (req *NetallianceAppletextensionlinkGetRequest) APIName() string {
	return "suning.netalliance.appletextensionlink.get"
}

// APIName 获取接口的名称
func (req *NetallianceAppletextensionlinkGetRequest) BizName() string {
	return "getAppletextensionlink"
}

func (req *NetallianceAppletextensionlinkGetRequest) Validate() {
	req.Form.ValidateRequired("productUrl")

}
