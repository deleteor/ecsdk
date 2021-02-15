package tb

// GetTime 获取淘宝系统当前时间
func (c *Client) GetTime(cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TimeGet)
	params := make(map[string]interface{})
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
