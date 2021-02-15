package tb

// GetOpenuid 获取授权账号对应的OpenUid
func (c *Client) GetOpenuid(cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(OpenuidGet)

	params := make(map[string]interface{})
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
