package responses

type NetallianceMorerecommendGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			GetMorerecommend struct {
				CommodityList []Commodityinfo `json:"commodityList"`
			} `json:"getMorerecommend"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}

type Commodityinfo struct {
	IsReserveCommodity int64 `json:"isReserveCommodity"`
	CommodityName  string `json:"commodityName"` //商品名称
	CommodityCode  string `json:"commodityCode"` //商品编码
	SupplierCode   string `json:"supplierCode"`  //供应商编码
	ImgVersion     string `json:"imgVersion"`  //图片版本号
	MonthlySales   int64 `json:"monthlySales"`  //月销量
	CouponAmount   int64 `json:"couponAmount"`  //券总数
	SnPrice        string `json:"snPrice"`  //易购价
	SellingPoint   string `json:"sellingPoint"` //亮点
	PriceType      string `json:"priceType"`  //价格类型
	PriceTypeCode  int64 `json:"priceTypeCode"`//0：正常价格，1、大聚惠 ， 2、抢购， 3、渠道专享， 4 限时促销 ，5爆款抢购， 6 掌上抢， 7特卖, 8 乐拼购    价格类型编码
	CommodityPrice string `json:"commodityPrice"` //商品价格
	CouponSpecialPrice string `json:"couponSpecialPrice"` //券后价
	CommissionPrice   string `json:"commissionPrice"`  //佣金金额
	CommissionRate    string `json:"commissionRate"`  //佣金比例
	CouponActiveId   string `json:"couponActiveId"`  //券活动id
	ActivitySecretKey string `json:"activitySecretKey"` //券活动秘钥
	CouponPrice   int64 `json:"couponPrice"` //券价格
	CommodityType int64 `json:"commodityType"` //商品类型
	Baoyou  int64 `json:"baoyou"` //包邮，0：不包邮，默认，1：包邮    是否包邮
	PgActionId int64 `json:"pgActionId"` //乐拼购活动ID
	PgNum  string `json:"pgNum"`  //成团人数
	PgPrice  string `json:"pgPrice"` //乐拼购价格
	IsOwnCommodity bool `json:"isOwnCommodity"`//0：是，1：否    是否自营
	SaleStatus  int64 `json:"saleStatus"` // 0:可售，1:无货，2:本地暂不销售    销售状态
	PicList   string `json:"picList"`  //图片集合

}

type CommodityPicList struct {
	CmmdtyUrl string `json:"cmmdtyUrl"` //图片URL
}

