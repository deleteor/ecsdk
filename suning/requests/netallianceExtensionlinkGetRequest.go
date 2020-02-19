package requests

type NetallianceExtensionlinkGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceExtensionlinkGetRequest() *NetallianceExtensionlinkGetRequest {
	req := &NetallianceExtensionlinkGetRequest{}
	req.Form = make(TopValues)
	return req
}

// SetProductUrl 商品的URL仅支持拼购和易购单品
func (req *NetallianceExtensionlinkGetRequest) SetProductUrl(productUrl string) {
	req.Form.Add("productUrl", productUrl)
}

// SetQuanUrl 券URL有时生成券推广，没有时生成商品推广
func (req *NetallianceExtensionlinkGetRequest) SetQuanUrl(quanUrl string) {
	req.Form.Add("quanUrl", quanUrl)
}

// SetPromotionId 这个推广只联盟前台申请的推广位，如果没有可以不填
func (req *NetallianceExtensionlinkGetRequest) SetPromotionId(promotionId string) {
	req.Form.Add("promotionId", promotionId)
}

// SetSubUser 子会员编码
func (req *NetallianceExtensionlinkGetRequest) SetSubUser(subUser string) {
	req.Form.Add("subUser", subUser)
}

// APIName 获取接口的名称
func (req *NetallianceExtensionlinkGetRequest) APIName() string {
	return "suning.netalliance.extensionlink.get"
}

// APIName 获取接口的名称
func (req *NetallianceExtensionlinkGetRequest) BizName() string {
	return "getExtensionlink"
}

func (req *NetallianceExtensionlinkGetRequest) Validate() {
	req.Form.ValidateRequired("productUrl")

}
