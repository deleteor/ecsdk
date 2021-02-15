package tb

// TbkItemInfoGetRequest 请求参数
type TbkItemInfoGetRequest struct {
	NumIids  string `json:"num_iids" form:"num_iids"`
	Platform int    `json:"platform" form:"platform"`
	IP       string `json:"ip" form:"ip"`
}

// GetTbkItemInfo 淘宝客-公用-淘宝客商品详情查询(简版)
func (c *Client) GetTbkItemInfo(in *TbkItemInfoGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkItemInfoGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
