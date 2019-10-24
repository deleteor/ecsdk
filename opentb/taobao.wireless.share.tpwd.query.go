package opentb

// WirelessShareTpwdQueryRequest 查询解析淘口令
type WirelessShareTpwdQueryRequest struct {
	// PasswordContent true	【天猫品牌号】，复制这条信息￥sMCl0Yra3Ae￥后打开手机淘宝	淘口令
	PasswordContent string `json:"password_content" form:"password_content"`
}

// QueryWirelessShareTpwd 查询解析淘口令
func (c *Client) QueryWirelessShareTpwd(in *WirelessShareTpwdQueryRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(WirelessShareTpwdQuery)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
