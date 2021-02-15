package requests

type NetallianceInverstmentcommodityQueryGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceInverstmentcommodityQueryGetRequest() *NetallianceInverstmentcommodityQueryGetRequest {
	req := &NetallianceInverstmentcommodityQueryGetRequest{}
	req.Form = make(TopValues)
	return req
}

// SetCategoryId 高佣类目ID 默认为第一个类目
func (req *NetallianceInverstmentcommodityQueryGetRequest) SetCategoryId(categoryId string) {
	req.Form.Add("categoryId", categoryId)
}

// SetPageIndex 页码 默认为0
func (req *NetallianceInverstmentcommodityQueryGetRequest) SetPageIndex(pageIndex int64) {
	req.Form.Add("pageIndex", pageIndex)
}

// SetPageSize 每页条数 默认10
func (req *NetallianceInverstmentcommodityQueryGetRequest) SetPageSize(pageSize int64) {
	req.Form.Add("size", pageSize)
}

// SetPicWidth 图片宽度 默认为200
func (req *NetallianceInverstmentcommodityQueryGetRequest) SetPicWidth(picWidth string) {
	req.Form.Add("picWidth", picWidth)
}

// SetPicHeight 图片高度 默认为200
func (req *NetallianceInverstmentcommodityQueryGetRequest) SetPicHeight(picHeight string) {
	req.Form.Add("picHeight", picHeight)
}

// SetCityCode 城市编码 默认为025
func (req *NetallianceInverstmentcommodityQueryGetRequest) SetCityCode(cityCode string) {
	req.Form.Add("cityCode", cityCode)
}

// APIName 获取接口的名称
func (req *NetallianceInverstmentcommodityQueryGetRequest) APIName() string {
	return "suning.netalliance.inverstmentcommodity.query"
}

// APIName 获取接口的名称
func (req *NetallianceInverstmentcommodityQueryGetRequest) BizName() string {
	return "queryInverstmentcommodity"
}

func (req *NetallianceInverstmentcommodityQueryGetRequest) Validate() {

}
