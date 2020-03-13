package requests

type NetallianceSearchcommodityQueryGetReuqest struct {
	TopRequest
}

//实例
func NewNetallianceSearchcommodityQueryGetReuqest() *NetallianceSearchcommodityQueryGetReuqest {
	req := &NetallianceSearchcommodityQueryGetReuqest{}
	req.Form = make(TopValues)
	return req
}

// SetPageIndex 页码 默认为1
func (req *NetallianceSearchcommodityQueryGetReuqest) SetPageIndex(pageIndex int64) {
	req.Form.Add("pageIndex", pageIndex)
}

// SetPageSize 每页条数 默认10
func (req *NetallianceSearchcommodityQueryGetReuqest) SetPageSize(pageSize int64) {
	req.Form.Add("size", pageSize)
}

// SetKeyword  关键字    与	销售目录ID 二选一
func (req *NetallianceSearchcommodityQueryGetReuqest) SetKeyword(keyword string) {
	req.Form.Add("keyword", keyword)
}

// SetSaleCategoryCode  销售目录ID   与	关键字 二选一
func (req *NetallianceSearchcommodityQueryGetReuqest) SetSaleCategoryCode(saleCategoryCode int64) {
	req.Form.Add("saleCategoryCode", saleCategoryCode)
}

// SetSuningService  是否苏宁自营 默认为空，1：是
func (req *NetallianceSearchcommodityQueryGetReuqest) SetSuningService(suningService int64) {
	req.Form.Add("suningService", suningService)
}

// SetPgSearch 是否拼购 默认为空 1：是
func (req *NetallianceSearchcommodityQueryGetReuqest) SetPgSearch(pgSearch int64) {
	req.Form.Add("pgSearch", pgSearch)
}

// SetStartPrice 开始价格
func (req *NetallianceSearchcommodityQueryGetReuqest) SetStartPrice(startPrice int64) {
	req.Form.Add("startPrice", startPrice)
}

// SetEndPrice 结束价格
func (req *NetallianceSearchcommodityQueryGetReuqest) SetEndPrice(endPrice int64) {
	req.Form.Add("endPrice", endPrice)
}

// SetSortType 排序规则 1：综合（默认） 2：销量由高到低 3：价格由高到低 4：价格由低到高
// 5：佣金比例由高到低 6：佣金金额由高到低 7：两个维度，佣金金额由高到低，销量由高到低
func (req *NetallianceSearchcommodityQueryGetReuqest) SetSortType(sortType int64) {
	req.Form.Add("sortType", sortType)
}

// SetBranch 1：减枝 2：不减枝 sortType=1（综合） 默认不剪枝 其他排序默认剪枝
func (req *NetallianceSearchcommodityQueryGetReuqest) SetBranch(branch int64) {
	req.Form.Add("branch", branch)
}

// SetCoupon 1：有券；其他：全部
func (req *NetallianceSearchcommodityQueryGetReuqest) SetCoupon(coupon int64) {
	req.Form.Add("coupon", coupon)
}

// SetPicWidth 图片宽度 默认为200
func (req *NetallianceSearchcommodityQueryGetReuqest) SetPicWidth(picWidth string) {
	req.Form.Add("picWidth", picWidth)
}

// SetPicHeight 图片高度 默认为200
func (req *NetallianceSearchcommodityQueryGetReuqest) SetPicHeight(picHeight string) {
	req.Form.Add("picHeight", picHeight)
}

// SetCityCode 城市编码 默认为025
func (req *NetallianceSearchcommodityQueryGetReuqest) SetCityCode(cityCode string) {
	req.Form.Add("cityCode", cityCode)
}

// APIName 获取接口的名称
func (req *NetallianceSearchcommodityQueryGetReuqest) APIName() string {
	return "suning.netalliance.searchcommodity.query"
}

// APIName 获取接口的名称
func (req *NetallianceSearchcommodityQueryGetReuqest) BizName() string {
	return "querySearchcommodity"
}

func (req *NetallianceSearchcommodityQueryGetReuqest) Validate() {

}
