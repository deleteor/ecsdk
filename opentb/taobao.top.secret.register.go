package opentb

// TopSecretRegisterRequest 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
type TopSecretRegisterRequest struct {
	// UserID false	12345	用户id，保证唯一
	UserID int `json:"user_id" form:"user_id"`
}

// RegisterTopSecret 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
func (c *Client) RegisterTopSecret(in *TopSecretRegisterRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TopSecretRegister)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
