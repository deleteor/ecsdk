package opentb

// TbkTpwdCreateRequest 淘宝客-公用-淘口令生成
type TbkTpwdCreateRequest struct {
	// UserID false	123	生成口令的淘宝用户ID
	UserID string `json:"user_id" form:"user_id"`
	// Text true	长度大于5个字符	口令弹框内容
	Text string `json:"text" form:"text"`
	// URL true	https://uland.taobao.com/	口令跳转目标页
	URL string `json:"url" form:"url"`
	// Logo false	https://uland.taobao.com/	口令弹框logoURL
	Logo string `json:"logo" form:"logo"`
	// Ext false	{}	扩展字段JSON格式
	Ext string `json:"ext" form:"ext"`
}

// CreateTbkTpwd 淘宝客-公用-淘口令生成
func (c *Client) CreateTbkTpwd(in *TbkTpwdCreateRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkTpwdCreate)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
