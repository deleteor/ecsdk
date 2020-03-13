package requests

type NetallianceCouponinfoQueryGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceCouponinfoQueryGetRequest() *NetallianceCouponinfoQueryGetRequest {
	req := &NetallianceCouponinfoQueryGetRequest{}
	req.Form = make(TopValues)
	return req
}


// SetQuanUrl  券URL
func (req *NetallianceCouponinfoQueryGetRequest) SetQuanUrl(quanUrl string) {
	req.Form.Add("quanUrl", quanUrl)
}




// APIName 获取接口的名称
func (req *NetallianceCouponinfoQueryGetRequest) APIName() string {
	return "suning.netalliance.couponinfo.query"
}
// APIName 获取接口的名称
func (req *NetallianceCouponinfoQueryGetRequest) BizName() string {
	return "queryCouponinfo"
}

func (req *NetallianceCouponinfoQueryGetRequest) Validate() {
	req.Form.ValidateRequired("quanUrl")

}