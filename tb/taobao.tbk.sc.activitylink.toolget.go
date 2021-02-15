package tb

// TbkScActivitylinkToolgetRequest 淘宝客-服务商-官方活动转链
type TbkScActivitylinkToolgetRequest struct {
	// AdzoneID true	123	推广位id，mm_xx_xx_xx pid三段式中的第三段
	AdzoneID int `json:"adzone_id" form:"adzone_id"`
	// SiteID true	456	推广位id，mm_xx_xx_xx pid三段式中的第二段
	SiteID int `json:"site_id" form:"site_id"`
	// Platform false	1	1：PC，2：无线，默认：１
	Platform int `json:"platform" form:"platform"`
	// UnionID false	demo	自定义输入串，英文和数字组成，长度不能大于12个字符，区分不同的推广渠道
	UnionID string `json:"union_id" form:"union_id"`
	// RelationID false	23	渠道关系ID，仅适用于渠道推广场景
	RelationID string `json:"relation_id" form:"relation_id"`
	// PromotionSceneID true	12345678	官方活动ID，从官方活动页获取
	PromotionSceneID int `json:"promotion_scene_id" form:"promotion_scene_id"`
}

// GetTbkScActivitylinkTool 淘宝客-服务商-官方活动转链
func (c *Client) GetTbkScActivitylinkTool(in *TbkScActivitylinkToolgetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScActivitylinkToolget)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
