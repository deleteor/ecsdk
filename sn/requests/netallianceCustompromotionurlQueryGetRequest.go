package requests

type NetallianceCustompromotionurlQueryGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceCustompromotionurlQueryGetRequest() *NetallianceCustompromotionurlQueryGetRequest {
	req := &NetallianceCustompromotionurlQueryGetRequest{}
	req.Form = make(TopValues)
	return req
}

// SetAdBookId 推广位id(在苏宁联盟会员前台已维护的推广位ID)
func (req *NetallianceCustompromotionurlQueryGetRequest) SetAdBookId(adBookId string) {
	req.Form.Add("adBookId", adBookId)
}

// SetVisitUrl 需要定制的链接(允许定制商品，店铺，频道及活动促销页面，其它页面暂不支持定制)
func (req *NetallianceCustompromotionurlQueryGetRequest) SetVisitUrl(visitUrl string) {
	req.Form.Add("visitUrl", visitUrl)
}


// APIName 获取接口的名称
func (req *NetallianceCustompromotionurlQueryGetRequest) APIName() string {
	return "suning.netalliance.custompromotionurl.query"
}

// APIName 获取接口的名称
func (req *NetallianceCustompromotionurlQueryGetRequest) BizName() string {
	return "queryCustompromotionurl"
}

func (req *NetallianceCustompromotionurlQueryGetRequest) Validate() {
	req.Form.ValidateRequired("adBookId")
	req.Form.ValidateRequired("visitUrl")
}
