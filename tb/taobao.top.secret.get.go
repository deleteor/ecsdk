package tb

// TopSecretGetRequest 获取TOP通道解密秘钥
type TopSecretGetRequest struct {
	// SecretVersion false	1	秘钥版本号
	SecretVersion int `json:"secret_version" form:"secret_version"`
	// RandomNum true	0ebbcccfee18d7ad1aebc5b875ffa906	伪随机数
	RandomNum string `json:"random_num" form:"random_num"`
	// CustomerUserID false	1111111	自定义用户id
	CustomerUserID int `json:"customer_user_id" form:"customer_user_id"`
}

// GetTopSecret 获取TOP通道解密秘钥
func (c *Client) GetTopSecret(in *TopSecretGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TopSecretGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
