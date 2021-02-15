package tb

// TopAuthTokenCreateRequest 获取Access Token 用户通过code换获取access_token，https only
type TopAuthTokenCreateRequest struct {
	// Code true	0_ebKlCPqc6OD8RBlB0DzfnpUg2	授权code，grantType==authorization_code 时需要
	Code string `json:"code" form:"code"`
	// UUID false	abc	与生成code的uuid配对
	UUID string `json:"uuid" form:"uuid"`
}

// CreateTopAuthToken 获取Access Token 用户通过code换获取access_token
func (c *Client) CreateTopAuthToken(in *TopAuthTokenCreateRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TopAuthTokenCreate)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
