package tb

// TbkScOptimusMaterialRequest  淘宝客-服务商-物料精选
type TbkScOptimusMaterialRequest struct {
	// PageSize false	20	页大小，默认20，1~100
	PageSize int `json:"page_size" form:"page_size"`
	// AdzoneID true	123	mm_xxx_xxx_xxx的第三位
	AdzoneID int `json:"adzone_id" form:"adzone_id"`
	// PageNo false	1	第几页，默认：1
	PageNo int `json:"page_no" form:"page_no"`
	// MaterialID false	123	官方的物料Id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)
	MaterialID int `json:"material_id" form:"material_id"`
	// SiteID true	111	mm_xxx_xxx_xxx的第二位
	SiteID int `json:"site_id" form:"site_id"`
	// DeviceType false	IMEI	智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	DeviceType string `json:"device_type" form:"device_type"`
	// DeviceEncrypt false	MD5	智能匹配-设备号加密类型：MD5
	DeviceEncrypt string `json:"device_encrypt" form:"device_encrypt"`
	// DeviceValue false	xxx	智能匹配-设备号加密后的值（MD5加密需32位小写）
	DeviceValue string `json:"device_value" form:"device_value"`
	// ContentID false	323	内容专用-内容详情ID
	ContentID int `json:"content_id" form:"content_id"`
	// ContentSource false	xxx	内容专用-内容渠道信息
	ContentSource string `json:"content_source" form:"content_source"`
	// ItemID false	33243	商品ID，用于相似商品推荐
	ItemID int `json:"item_id" form:"item_id"`
}

// GetTbkScOptimusMaterial 淘宝客-服务商-物料精选
func (c *Client) GetTbkScOptimusMaterial(in *TbkScOptimusMaterialRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScOptimusMaterial)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
