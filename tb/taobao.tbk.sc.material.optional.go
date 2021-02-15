package tb

// TbkScMaterialOptionalRequest 淘宝客-服务商-物料搜索
type TbkScMaterialOptionalRequest struct {
	// StartDsr	false	10	商品筛选(特定媒体支持)-店铺dsr评分。筛选大于等于当前设置的店铺dsr评分的商品0-50000之间
	StartDsr int `json:"start_dsr" form:"start_dsr"`
	// PageSize	false	20	页大小，默认20，1~100
	PageSize int `json:"page_size" form:"page_size"`
	// PageNo	false	1	第几页，默认：１
	PageNo int `json:"page_no" form:"page_no"`
	// Platform	false	1	链接形式：1：PC，2：无线，默认：１
	Platform int `json:"platform" form:"platform"`
	// EndTkRate	false	1234	商品筛选-淘客佣金比率上限。如：1234表示12.34%
	EndTkRate int `json:"end_tk_rate" form:"end_tk_rate"`
	// StartTkRate	false	1234	商品筛选-淘客佣金比率下限。如：1234表示12.34%
	StartTkRate int `json:"start_tk_rate" form:"start_tk_rate"`
	// EndPrice	false	10	商品筛选-折扣价范围上限。单位：元
	EndPrice int `json:"end_price" form:"end_price"`
	// StartPrice	false	10	商品筛选-折扣价范围上限。单位：元
	StartPrice int `json:"start_price" form:"start_price"`
	// IsOverseas	false	false	商品筛选-是否海外商品。true表示属于海外商品，false或不设置表示不限
	IsOverseas bool `json:"is_overseas" form:"is_overseas"`
	// IsTmall	false	false	商品筛选-是否天猫商品。true表示属于天猫商品，false或不设置表示不限
	IsTmall bool `json:"is_tmall" form:"is_tmall"`
	// Sort	false	tk_rate_des	排序_des（降序），排序_asc（升序），销量（total_sales），淘客佣金比率（tk_rate）， 累计推广量（tk_total_sales），总支出佣金（tk_total_commi），价格（price）
	Sort string `json:"sort" form:"sort"`
	// Itemloc	false	杭州	商品筛选-所在地
	Itemloc string `json:"itemloc" form:"itemloc"`
	// Cat	false	16,18	商品筛选-后台类目ID。用,分割，最大10个，该ID可以通过taobao.itemcats.get接口获取到
	Cat string `json:"cat" form:"cat"`
	// Q	false	女装	查询的关键词
	Q string `json:"q" form:"q"`
	// AdzoneID	true	12345678	mm_xxx_xxx_12345678三段式的最后一段数字
	AdzoneID int `json:"adzone_id" form:"adzone_id"`
	// SiteID	true	22	mm_xxx_22_xxx三段式的第二段数字
	SiteID int `json:"site_id" form:"site_id"`
	// MaterialID	false	2836	物料id(详细物料id见：https://tbk.bbs.taobao.com/detail.html?appId=45301&postId=8576096)。不传时默认为2836
	MaterialID int `json:"material_id" form:"material_id"`
	// HasCoupon	false	false	优惠券筛选-是否有优惠券。true表示该商品有优惠券，false或不设置表示不限
	HasCoupon bool `json:"has_coupon" form:"has_coupon"`
	// IP	false	13.2.33.4	ip参数影响邮费获取，如果不传或者传入不准确，邮费无法精准提供
	IP string `json:"ip" form:"ip"`
	// IncludeRfdRate	false	true	商品筛选(特定媒体支持)-退款率是否低于行业均值。True表示大于等于，false或不设置表示不限
	IncludeRfdRate bool `json:"include_rfd_rate" form:"include_rfd_rate"`
	// IncludeGoodRate	false	true	商品筛选-好评率是否高于行业均值。True表示大于等于，false或不设置表示不限
	IncludeGoodRate bool `json:"include_good_rate" form:"include_good_rate"`
	// IncludePayRate30	false	true	商品筛选(特定媒体支持)-成交转化是否高于行业均值。True表示大于等于，false或不设置表示不限
	IncludePayRate30 bool `json:"include_pay_rate_30" form:"include_pay_rate_30"`
	// NeedPrepay	false	true	商品筛选-是否加入消费者保障。true表示加入，false或不设置表示不限
	NeedPrepay bool `json:"need_prepay" form:"need_prepay"`
	// NeedFreeShipment	false	true	商品筛选-是否包邮。true表示包邮，false或不设置表示不限
	NeedFreeShipment bool `json:"need_free_shipment" form:"need_free_shipment"`
	// NpxLevel	false	2	商品筛选-牛皮癣程度。取值：1不限，2无，3轻微
	NpxLevel int `json:"npx_level" form:"npx_level"`
	// EndKaTkRate	false	1234	商品筛选-KA媒体淘客佣金比率上限。如：1234表示12.34%
	EndKaTkRate int `json:"end_ka_tk_rate" form:"end_ka_tk_rate"`
	// StartKaTkRate	false	1234	商品筛选-KA媒体淘客佣金比率下限。如：1234表示12.34%
	StartKaTkRate int `json:"start_ka_tk_rate" form:"start_ka_tk_rate"`
	// DeviceValue	false	xxx	智能匹配-设备号加密后的值（MD5加密需32位小写）
	DeviceValue string `json:"device_value" form:"device_value"`
	// DeviceEncrypt	false	MD5	智能匹配-设备号加密类型：MD5
	DeviceEncrypt string `json:"device_encrypt" form:"device_encrypt"`
	// DeviceType	false	IMEI	智能匹配-设备号类型：IMEI，或者IDFA，或者UTDID（UTDID不支持MD5加密），或者OAID
	DeviceType string `json:"device_type" form:"device_type"`
	// LockRateEndTime	false	1567440000000	锁佣结束时间
	LockRateEndTime int `json:"lock_rate_end_time" form:"lock_rate_end_time"`
	// LockRateStartTime	false	1567440000000	锁佣开始时间
	LockRateStartTime int `json:"lock_rate_start_time" form:"lock_rate_start_time"`
}

// GetTbkScMaterialOptional 淘宝客-服务商-物料搜索
func (c *Client) GetTbkScMaterialOptional(in *TbkScMaterialOptionalRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScMaterialOptional)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
