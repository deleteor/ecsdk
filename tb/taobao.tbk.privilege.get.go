package tb

// TbkPrivilegeGetRequest 淘宝客-服务商-单品券高效转链
type TbkPrivilegeGetRequest struct {
	// ItemID false	123	淘客商品id
	ItemID int `json:"item_id" form:"item_id"`
	// AdzoneID true	123	推广位id，mm_xx_xx_xx pid三段式中的第三段
	AdzoneID int `json:"adzone_id" form:"adzone_id"`
	// Platform false	1	1：PC，2：无线，默认：１
	Platform int `json:"platform" form:"platform"`
	// SiteID true	1	备案的网站id, mm_xx_xx_xx pid三段式中的第二段
	SiteID int `json:"site_id" form:"site_id"`
	// Me false	m%3D2%26s%3D94BYV45NHwgcQipKwQzePOeEDrYVVa64LKpWJ%2Bin0XLjf2vlNIV67pL2V8ikcqW7FfrEfJ4hp2q5rze35H1YEElKMSinFVD02hfsaefZn7H4%2Ff3V	营销计划链接中的me参数
	Me string `json:"me" form:"me"`
	// RelationID false	23223	渠道关系ID，仅适用于渠道推广场景
	RelationID string `json:"relation_id" form:"relation_id"`
}

// GetTbkPrivilege 淘宝客-服务商-单品券高效转链
func (c *Client) GetTbkPrivilege(in *TbkPrivilegeGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkPrivilegeGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
