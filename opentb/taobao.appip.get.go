package opentb

// GetAppIP 获取ISV发起请求服务器IP
func (c *Client) GetAppIP(cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(AppIPGet)
	params := make(map[string]interface{})
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
