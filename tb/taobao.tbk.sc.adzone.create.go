package tb

// TbkScAdzoneCreateRequest 淘宝客-服务商-创建推广者位
type TbkScAdzoneCreateRequest struct {
	// SiteID true	123456	网站ID
	SiteID int `json:"site_id" form:"site_id"`
	// AdzoneName true	广告位	广告位名称，最大长度64字符
	AdzoneName string `json:"adzone_name" form:"adzone_name"`
}

// CreateTbkScAdzone 淘宝客-服务商-创建推广者位
func (c *Client) CreateTbkScAdzone(in *TbkScAdzoneCreateRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScAdzoneCreate)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
