package requests

type NetallianceRecommendcommodityQueryGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceRecommendcommodityQueryGetRequest() *NetallianceRecommendcommodityQueryGetRequest {
	req := &NetallianceRecommendcommodityQueryGetRequest{}
	req.Form = make(TopValues)
	return req
}


// SetPageIndex  页码 默认0
func (req *NetallianceRecommendcommodityQueryGetRequest) SetPageIndex(pageIndex int64) {
	req.Form.Add("pageIndex", pageIndex)
}

// SetPageSize  每页数量，默认10
func (req *NetallianceRecommendcommodityQueryGetRequest) SetPageSize(size int64) {
	req.Form.Add("size", size)
}


// SetPicWidth 图片宽度 默认为200
func (req *NetallianceRecommendcommodityQueryGetRequest) SetPicWidth(picWidth string) {
	req.Form.Add("picWidth", picWidth)
}

// SetPicHeight 图片高度 默认为200
func (req *NetallianceRecommendcommodityQueryGetRequest) SetPicHeight(picHeight string) {
	req.Form.Add("picHeight", picHeight)
}

// SetCityCode 城市编码 默认为025
func (req *NetallianceRecommendcommodityQueryGetRequest) SetCityCode(cityCode string) {
	req.Form.Add("cityCode", cityCode)
}


// APIName 获取接口的名称
func (req *NetallianceRecommendcommodityQueryGetRequest) APIName() string {
	return "suning.netalliance.recommendcommodity.query"
}
// APIName 获取接口的名称
func (req *NetallianceRecommendcommodityQueryGetRequest) BizName() string {
	return "queryRecommendcommodity"
}

func (req *NetallianceRecommendcommodityQueryGetRequest) Validate() {

}