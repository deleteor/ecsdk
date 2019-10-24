package opentb

// GetTopIpout 获取开放平台出口IP段
func (c *Client) GetTopIpout(cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TopIpoutGet)
	params := make(map[string]interface{})
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
