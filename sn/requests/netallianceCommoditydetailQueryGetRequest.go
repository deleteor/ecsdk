package requests

type NetallianceCommoditydetailQueryGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceCommoditydetailQueryGetRequest() *NetallianceCommoditydetailQueryGetRequest {
	req := &NetallianceCommoditydetailQueryGetRequest{}
	req.Form = make(TopValues)
	return req
}

// SetCommodityStr 设置格式 商品编码1-供应商编码1_商品编码2-供应商编码2...
// 商品编码取有效位，供应商编码左补零至10位 最大查询10个商品
func (req *NetallianceCommoditydetailQueryGetRequest) SetCommodityStr(commodityStr string) {
	req.Form.Add("commodityStr", commodityStr)
}

// SetPicWidth 图片宽度 默认为200
func (req *NetallianceCommoditydetailQueryGetRequest) SetPicWidth(picWidth string) {
	req.Form.Add("picWidth", picWidth)
}

// SetPicHeight 图片高度 默认为200
func (req *NetallianceCommoditydetailQueryGetRequest) SetPicHeight(picHeight string) {
	req.Form.Add("picHeight", picHeight)
}

// SetCityCode 城市编码 默认为025
func (req *NetallianceCommoditydetailQueryGetRequest) SetCityCode(cityCode string) {
	req.Form.Add("cityCode", cityCode)
}

// APIName 获取接口的名称
func (req *NetallianceCommoditydetailQueryGetRequest) APIName() string {
	return "suning.netalliance.commoditydetail.query"
}

// APIName 获取接口的名称
func (req *NetallianceCommoditydetailQueryGetRequest) BizName() string {
	return "queryCommoditydetail"
}

func (req *NetallianceCommoditydetailQueryGetRequest) Validate() {
	req.Form.ValidateRequired("commodityStr")

}
