package tb

// TbkScOrderDetailsGetRequest 淘宝客-服务商-所有订单查询
type TbkScOrderDetailsGetRequest struct {
	// QueryType false	1	查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询
	QueryType int `json:"query_type" form:"query_type"`
	// PositionIndex false	2222_334666	位点，除第一页之外，都需要传递；前端原样返回。
	PositionIndex string `json:"position_index" form:"position_index"`
	// PageSize false	20	页大小，默认20，1~100
	PageSize int `json:"page_size" form:"page_size"`
	// MemberType false	2	推广者角色类型,2:二方，3:三方，不传，表示所有角色
	MemberType int `json:"member_type" form:"member_type"`
	// TkStatus false	12	淘客订单状态，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态
	TkStatus int `json:"tk_status" form:"tk_status"`
	// EndTime true	2019-04-23 12:28:22	订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！
	EndTime string `json:"end_time" form:"end_time"`
	// StartTime true	2019-04-05 12:18:22	订单查询开始时间
	StartTime string `json:"start_time" form:"start_time"`
	// JumpType false	1	跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页
	JumpType int `json:"jump_type" form:"jump_type"`
	// PageNo false	1	第几页，默认1，1~100
	PageNo int `json:"page_no" form:"page_no"`
	// OrderScene false	1	场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1
	OrderScene int `json:"order_scene" form:"order_scene"`
}

// GetTbkScOrderDetails 淘宝客-服务商-所有订单查询
func (c *Client) GetTbkScOrderDetails(in *TbkScOrderDetailsGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScOrderDetailsGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
