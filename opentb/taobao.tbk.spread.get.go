package opentb

// TbkSpreadGetRequest 淘宝客-公用-长链转短链
type TbkSpreadGetRequest struct {
	// Requests true		请求列表，内部包含多个url
	Requests []*TbkSpreadURL
}

// TbkSpreadURL 需要转换的URL
type TbkSpreadURL struct {
	// URL true	http://temai.taobao.com	原始url, 只支持uland.taobao.com，s.click.taobao.com， ai.taobao.com，temai.taobao.com的域名转换，否则判错
	URL string `json:"url" form:"url"`
}

// GetTbkSpread 淘宝客-公用-长链转短链
func (c *Client) GetTbkSpread(in *TbkSpreadGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkSpreadGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
