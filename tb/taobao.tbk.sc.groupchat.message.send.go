package tb

// TbkScGroupchatMessageSendRequest 淘宝客-服务商-手淘群发单
type TbkScGroupchatMessageSendRequest struct {
	// PID true	mm_0_0	淘客pid
	PID string `json:"pid" form:"pid"`
	// GroupIDs true	v134,v234	消息发送群组id列表，逗号分隔
	GroupIDs string `json:"group_ids" form:"group_ids"`
	// ItemID false	12334	商品id
	ItemID int `json:"item_id" form:"item_id"`
	// ActivityID false	kf123	券id
	ActivityID string `json:"activity_id" form:"activity_id"`
	// Me false	234df	加密e参数
	Me string `json:"me" form:"me"`
	// Dx false	1	强制定向，通用佣金结算
	Dx string `json:"dx" form:"dx"`
	// Text false	你好请购买	消息文本
	Text string `json:"text" form:"text"`
}

// SendTbkScGroupchatMessage 淘宝客-服务商-手淘群发单
func (c *Client) SendTbkScGroupchatMessage(in *TbkScGroupchatMessageSendRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScGroupchatMessageSend)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
