package responses

type ErrorResponse struct {
	Error ErrorObject `json:"sn_error,omitempty"`
}

type ErrorObject struct {
	SubCode   string `json:"error_code,omitempty"`
	SubMsg    string `json:"error_msg,omitempty"`
}


type QueryCommoditydetail struct {
	CommodityInfo CommodityInfo `json:"commodityInfo"`
	PgInfo        PgInfo        `json:"pgInfo"`
	CategoryInfo  CategoryInfo  `json:"categoryInfo"`
	CouponInfo    CouponInfo    `json:"couponInfo"`
}


//优惠券信息
type CouponInfo struct {
	CouponUrl           string `json:"couponUrl"`           //券链接URL
	ActivityId          string `json:"activityId"`          //券活动编码
	ActivitySecretKey   string `json:"activitySecretKey"`   //券秘钥
	CouponValue         string `json:"couponValue"`         //券面额
	CouponCount         string `json:"couponCount"`         //券总数
	CouponStartTime     string `json:"couponStartTime"`     //券领取开始时间
	CouponEndTime       string `json:"couponEndTime"`       //券领取结束时间
	StartTime           string `json:"startTime"`           //券使用开始时间
	EndTime             string `json:"endTime"`             //券使用结束时间
	BounsLimit          string `json:"bounsLimit"`          //基数数值
	ActivityDescription string `json:"activityDescription"` //活动描述
}

//图片结合
type PictureUrlList struct {
	LocationId int64  `json:"locationId"`
	PicUrl     string `json:"picUrl"`
}
//分类类目
type CategoryInfo struct {
	FirstSaleCategoryId        string `json:"firstSaleCategoryId"`        //一级销售目录ID
	FirstSaleCategoryName      string `json:"firstSaleCategoryName"`      //一级销售目录名称
	SecondSaleCategoryId       string `json:"secondSaleCategoryId"`       //二级销售目录ID
	SecondSaleCategoryName     string `json:"secondSaleCategoryName"`     //二级销售目录名称
	ThirdSaleCategoryId        string `json:"thirdSaleCategoryId"`        //三级销售目录ID
	ThirdSaleCategoryName      string `json:"thirdSaleCategoryName"`      //三级销售目录名称
	FirstPurchaseCategoryId    string `json:"firstPurchaseCategoryId"`    //一级采购目录ID
	FirstPurchaseCategoryName  string `json:"firstPurchaseCategoryName"`  //一级采购目录名称
	SecondPurchaseCategoryId   string `json:"secondPurchaseCategoryId"`   //二级采购目录ID
	SecondPurchaseCategoryName string `json:"secondPurchaseCategoryName"` //二级采购目录名称示例	二级采购目录名称
	ThirdPurchaseCategoryId    string `json:"thirdPurchaseCategoryId"`    //三级采购目录ID
	ThirdPurchaseCategoryName  string `json:"thirdPurchaseCategoryName"`  //三级采购目录名称示例	三级采购目录名称
	GoodsGroupCategoryId       string `json:"goodsGroupCategoryId"`       //商品组目录ID
	GoodsGroupCategoryName     string `json:"goodsGroupCategoryName"`     //商品组目录名称
}
//拼团信息
type PgInfo struct {
	PgNum      string `json:"pgNum"`
	PgPrice    string `json:"pgPrice"`
	PgUrl      string `json:"pgUrl"`
	PgActionId string `json:"pgActionId"`
}
//商品信息
type CommodityInfo struct {
	CommodityName  string           `json:"commodityName"`
	CommodityCode  string           `json:"commodityCode"`
	SupplierCode   string           `json:"supplierCode"`
	SupplierName   string           `json:"supplierName"`
	PictureUrl     []PictureUrlList `json:"pictureUrl"`
	SellingPoint   string           `json:"sellingPoint"`
	MonthSales     int64            `json:"monthSales"`
	SnPrice        string           `json:"snPrice"`
	CommodityPrice string           `json:"commodityPrice"`
	CommodityType  int64            `json:"commodityType"`
	PriceType      string           `json:"priceType"`
	PriceTypeCode  int64            `json:"priceTypeCode"`
	Baoyou         int64            `json:"baoyou"`
	Rrate          string           `json:"rate"`
	SaleStatus     int64            `json:"saleStatus"`
}
