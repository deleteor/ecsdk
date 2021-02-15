package tb

// GetTbkScGroupchat 获取ISV发起请求服务器IP
func (c *Client) GetTbkScGroupchat(cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScGroupchatGet)
	params := make(map[string]interface{})
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
