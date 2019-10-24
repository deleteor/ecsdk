package opentb

// TbkScPublisherInfoSaveRequest 淘宝客-公用-私域用户备案
type TbkScPublisherInfoSaveRequest struct {
	// RelationFrom false	1	渠道备案 - 来源，取链接的来源
	RelationFrom string `json:"relation_from" form:"relation_from"`
	// OfflineScene false	1	渠道备案 - 线下场景信息，1 - 门店，2- 学校，3 - 工厂，4 - 其他
	OfflineScene string `json:"offline_scene" form:"offline_scene"`
	// OnlineScene false	1	渠道备案 - 线上场景信息，1 - 微信群，2- QQ群，3 - 其他
	OnlineScene string `json:"online_scene" form:"online_scene"`
	// InviterCode true	xxx	渠道备案 - 淘宝客邀请渠道的邀请码
	InviterCode string `json:"inviter_code" form:"inviter_code"`
	// InfoType true	1	类型，必选 默认为1:
	InfoType int `json:"info_type" form:"info_type"`
	// Note false	小蜜蜂	媒体侧渠道备注
	Note string `json:"note" form:"note"`
	// RegisterInfo false
	RegisterInfo string `json:"register_info" form:"register_info"`
}

// SaveTbkScPublisherInfo 淘宝客-公用-私域用户备案
func (c *Client) SaveTbkScPublisherInfo(in *TbkScPublisherInfoSaveRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScPublisherInfoSave)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
