package tb

// GetHTTPSDns Top Dns 配置
func (c *Client) GetHTTPSDns(cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(HTTPDnsGet)

	params := make(map[string]interface{})
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
