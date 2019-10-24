package opentb

// TopAuthTokenRefreshRequest 刷新Access Token 根据refresh_token重新生成token
type TopAuthTokenRefreshRequest struct {
	// RefreshToken  true	50003401926fExdtvhBdiQ5PyVllfKjYEVWngx129d2b42DxpjypaTqrxxaWVSCIuIi1	grantType==refresh_token 时需要
	RefreshToken string `json:"refresh_token" form:"refresh_token"`
}

// RefreshTopAuthToken 淘宝客-公用-淘宝客商品详情查询(简版)
func (c *Client) RefreshTopAuthToken(in *TopAuthTokenRefreshRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TopAuthTokenRefresh)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
