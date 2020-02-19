package requests

type NetallianceInverstmentcategoryidQueryGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceInverstmentcategoryidQueryGetRequest() *NetallianceInverstmentcategoryidQueryGetRequest {
	req := &NetallianceInverstmentcategoryidQueryGetRequest{}
	req.Form = make(TopValues)
	return req
}

// APIName 获取接口的名称
func (req *NetallianceInverstmentcategoryidQueryGetRequest) APIName() string {
	return "suning.netalliance.inverstmentcategoryid.query"
}

// APIName 获取接口的名称
func (req *NetallianceInverstmentcategoryidQueryGetRequest) BizName() string {
	return "queryInverstmentcategoryid"
}

func (req *NetallianceInverstmentcategoryidQueryGetRequest) Validate() {

}
