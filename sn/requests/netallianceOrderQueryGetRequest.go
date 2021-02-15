package requests

type NetallianceOrderQueryGetRequest struct {
	TopRequest
}

//实例
func NewNetallianceOrderQueryGetRequest() *NetallianceOrderQueryGetRequest {
	req := &NetallianceOrderQueryGetRequest{}
	req.Form = make(TopValues)
	return req
}

// SetPageNo 页码。默认值：1。页码和每页条数成对出现
func (req *NetallianceOrderQueryGetRequest) SetPageNo(pageNo int64) {
	req.Form.Add("pageNo", pageNo)
}

// SetPageSize 每页条数。最大值：50，默认值：10，页码和每页条数成对出现
func (req *NetallianceOrderQueryGetRequest) SetPageSize(pageSize int64) {
	req.Form.Add("pageSize", pageSize)
}

//SetStartTime 查询开始时间。格式:yyyy-MM-dd HH:mm:ss，
// 若输入则开始时间和结束时间必须同时出现，开始时间和结束时间间隔不大于1天
func (req *NetallianceOrderQueryGetRequest) SetStartTime(startTime string) {
	req.Form.Add("startTime", startTime)
}

// SetCityCode 城市编码 默认为025
func (req *NetallianceOrderQueryGetRequest) SetEndTime(endTime string) {
	req.Form.Add("endTime", endTime)
}

// SetOrderLineStatus 订单行项目状态（0：全部状态；1：等待付款；2：支付完成；3：退款；4：订单已取消；5：确认收货）
func (req *NetallianceOrderQueryGetRequest) SetOrderLineStatus(orderLineStatus int64) {
	req.Form.Add("orderLineStatus", orderLineStatus)
}

// APIName 获取接口的名称
func (req *NetallianceOrderQueryGetRequest) APIName() string {
	return "suning.netalliance.order.query"
}

// APIName 获取接口的名称
func (req *NetallianceOrderQueryGetRequest) BizName() string {
	return "queryOrder"
}

func (req *NetallianceOrderQueryGetRequest) Validate() {
	req.Form.ValidateRequired("pageNo")
	req.Form.ValidateRequired("pageSize")
	req.Form.ValidateRequired("startTime")
	req.Form.ValidateRequired("endTime")

}
