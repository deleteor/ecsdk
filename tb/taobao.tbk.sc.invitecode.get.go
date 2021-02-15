package tb

// TbkScInvitecodeGetRequest 淘宝客-公用-私域用户邀请码生成
type TbkScInvitecodeGetRequest struct {
	// RelationID false	11	渠道关系ID
	RelationID int `json:"relation_id" form:"relation_id"`
	// RelationApp true	common	渠道推广的物料类型
	RelationApp string `json:"relation_app" form:"relation_app"`
	// CodeType true	1	邀请码类型，1 - 渠道邀请，2 - 渠道裂变，3 -会员邀请
	CodeType int `json:"code_type" form:"code_type"`
}

// GetTbkScInvitecode 淘宝客-公用-私域用户邀请码生成
func (c *Client) GetTbkScInvitecode(in *TbkScInvitecodeGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScInvitecodeGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
