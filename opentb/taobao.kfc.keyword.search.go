package opentb

// KeywordSearchRequest 请求参数
type KeywordSearchRequest struct {
	// Nick	 	false	淘宝会员名	发布信息的淘宝会员名，可以不传
	Nick string `json:"nick" form:"nick"`
	// Content 	true	文本信息	需要过滤的文本信息
	Content string `json:"content" form:"content"`
	// Apply 	false	taobao_auction.title	应用点，分为一级应用点、二级应用点。
	Apply string `json:"apply" form:"apply"`
}

// SearchKFCKeyword 关键词过滤匹配
func (c *Client) SearchKFCKeyword(in *KeywordSearchRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(KFCKeywordSearch)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
