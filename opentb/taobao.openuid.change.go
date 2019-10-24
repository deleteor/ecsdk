package opentb

// OpenuidChangeRequest 请求参数
type OpenuidChangeRequest struct {
	// OpenUID	true	123123	openUid
	OpenUID string `json:"open_uid" form:"open_uid"`
	// TargetAppKey	true	12011111	转换到的appkey
	TargetAppKey string `json:"target_app_key" form:"target_app_key"`
}

// ChangeOpenuid 淘宝客-公用-淘宝客商品详情查询(简版)
func (c *Client) ChangeOpenuid(in *OpenuidChangeRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(OpenuidChange)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
