package requests

import "fmt"

type NetallianceMorerecommendGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceMorerecommendGetRequest() *NetallianceMorerecommendGetRequest {
	req := &NetallianceMorerecommendGetRequest{}
	req.Form = make(TopValues)
	return req
}


// SetCommodityCode  商品编码
func (req *NetallianceMorerecommendGetRequest) SetCommodityCode(commodityCode string) {
	req.Form.Add("commodityCode", commodityCode)
}

// SetSupplierCode  商家编码
func (req *NetallianceMorerecommendGetRequest) SetSupplierCode(supplierCode string) {

	supplierCode = fmt.Sprintf("%010s", supplierCode)
	req.Form.Add("supplierCode", supplierCode)
}


// SetPicWidth 图片宽度 默认为200
func (req *NetallianceMorerecommendGetRequest) SetPicWidth(picWidth string) {
	req.Form.Add("picWidth", picWidth)
}

// SetPicType 图片类型 0方图，1长图，2竖图,默认为0
func (req *NetallianceMorerecommendGetRequest) SetPicType(picType string) {
	req.Form.Add("picType", picType)
}

// SetPicLocation 图片位置，默认值2
func (req *NetallianceMorerecommendGetRequest) SetPicLocation(picLocation string) {
	req.Form.Add("picLocation", picLocation)
}

// SetPicHeight 图片高度 默认为200
func (req *NetallianceMorerecommendGetRequest) SetPicHeight(picHeight string) {
	req.Form.Add("picHeight", picHeight)
}

// SetCityCode 城市编码 默认为025
func (req *NetallianceMorerecommendGetRequest) SetCityCode(cityCode string) {
	req.Form.Add("cityCode", cityCode)
}

// APIName 获取接口的名称
func (req *NetallianceMorerecommendGetRequest) APIName() string {
	return "suning.netalliance.morerecommend.get"
}
// APIName 获取接口的名称
func (req *NetallianceMorerecommendGetRequest) BizName() string {
	return "getMorerecommend"
}

func (req *NetallianceMorerecommendGetRequest) Validate() {
	req.Form.ValidateRequired("commodityCode")
	req.Form.ValidateRequired("supplierCode")
}