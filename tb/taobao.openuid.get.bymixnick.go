package tb

// OpenuidGetBymixnickRequest 请求参数
type OpenuidGetBymixnickRequest struct {
	// MixNick true	淘012xBbgHy7GJqSgfChglzpj82DEMds1FAQI7fgfrP3PMg=	无线类应用获取到的混淆的nick
	MixNick string `json:"mix_nick" form:"mix_nick"`
}

// GetOpenuidGetBymixnick 通过mixnick转换openuid
func (c *Client) GetOpenuidGetBymixnick(in *OpenuidGetBymixnickRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(OpenuidGetBymixnick)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
