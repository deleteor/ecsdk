package vip

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// UnionGoodsServiceName 服务名
	UnionGoodsServiceName string = "com.vip.adp.api.open.service.UnionGoodsService"
	// UnionGoodsVersion 版本号
	UnionGoodsVersion string = "1.0.0"
	// UnionGoodsMethodName 获取订单信息列表
	UnionGoodsMethodName string = "goodsList"
)

// GoodsInfoRequest 获取商品信息列表 应用级请求参数
type GoodsInfoRequest struct {
	// page Integer 页码
	Page        int    `json:"page"`
	ChannelType int    `json:"channelType"`
	RequestID   string `json:"requestId"`
}

// QueryGoodsList 获取商品信息列表
func (c *Client) QueryGoodsList(in *GoodsInfoRequest) (*GoodsInfoResponse, error) {

	c.SetMethod(UnionGoodsMethodName)
	c.SetServiceName(UnionGoodsServiceName)
	c.SetVersion(UnionGoodsVersion)

	params := GetParams(in)
	mapParams := make(map[string]interface{})
	mapParams["request"] = params
	result, err := c.DoRequest(mapParams)
	if err != nil {
		return nil, err
	}

	// fmt.Println(result)
	out := &UnionGoodsServiceGoodsListResponse{}
	outByte, _ := json.Marshal(result)
	json.Unmarshal(outByte, out)
	if out.ReturnCode != 0 {
		er := fmt.Sprintf("获取商品信息失败 code=%d\n", out.ReturnCode)
		return nil, errors.New(er)
	}
	return &out.Result, nil
}

// UnionGoodsServiceGoodsListResponse 商品信息相应参数
type UnionGoodsServiceGoodsListResponse struct {
	ReturnCode int               `json:"returnCode"`
	Result     GoodsInfoResponse `json:"result"`
}

// GoodsInfoResponse 商品信息相应参数
type GoodsInfoResponse struct {
	// goodsInfoList	List<GoodsInfo>	精选商品列表
	GoodsInfoList []GoodsInfo `json:"goodsInfoList"`
	// total	Integer	总数
	Total int `json:"total"`
	// sortFields	List<SortField>	支持的排序字段,为空时仅支持默认排序
	SortFields []struct {
		// fieldName	string	排序字段
		FieldName string `json:"fieldName"`
		// fieldDesc	string	字段说明
		FieldDesc string `json:"fieldDesc"`
	} `json:"sortFields"`
	// page	Integer	当前页码
	Page int `json:"page"`
	// pageSize	Integer	分页大小
	PageSize int `json:"pageSize"`
}

// GoodsInfo 精选商品列表
type GoodsInfo struct {
	// goodsId	String	商品id
	GoodsID string `json:"goodsId"`
	// goodsName	String	商品名称
	GoodsName string `json:"goodsName"`
	// goodsDesc	String	商品描述,字段暂不输出
	GoodsDesc string `json:"goodsDesc"`
	// destUrl	String	商品落地页
	DestURL string `json:"destUrl"`
	// goodsThumbUrl	String	商品缩略图
	GoodsThumbURL string `json:"goodsThumbUrl"`
	// goodsCarouselPictures	List<String>	商品轮播图：根据商品id查询时返回，商品列表不返回
	GoodsCarouselPictures []string `json:"goodsCarouselPictures"`
	// goodsMainPicture	String	商品主图
	GoodsMainPicture string `json:"goodsMainPicture"`
	// categoryId	Long	商品三级分类id
	CategoryID int64 `json:"categoryId"`
	// categoryName	String	商品三级分类
	CategoryName string `json:"categoryName"`
	// sourceType	Integer	商品类型：0-自营，1-MP
	SourceType int `json:"sourceType"`
	// marketPrice	String	市场价（元）
	MarketPrice string `json:"marketPrice"`
	// vipPrice	String	唯品价（元）
	VipPrice string `json:"vipPrice"`
	// commissionRate	String	佣金比例（%）
	CommissionRate string `json:"commissionRate"`
	// commission	String	佣金金额（元）
	Commission string `json:"commission"`
	// discount	String	折扣:唯品价/市场价 保留两位小数字符串
	Discount string `json:"discount"`
	// goodsDetailPictures	List<String>	商品详情图片：根据商品id查询商品信息时返回，商品列表不返回
	GoodsDetailPictures []string `json:"goodsDetailPictures"`
	// cat1stId	Long	商品一级分类id
	Cat1stID int64 `json:"cat1stId"`
	// cat1stName	String	商品一级分类名称
	Cat1stName string `json:"cat1stName"`
	// cat2ndId	Long	商品二级分类id
	Cat2ndID int64 `json:"cat2ndId"`
	// cat2ndName	String	商品二级分类名称
	Cat2ndName string `json:"cat2ndName"`
	// brandStoreSn	String	商品品牌sn
	BrandStoreSn string `json:"brandStoreSn"`
	// brandName	String	商品品牌名称
	BrandName string `json:"brandName"`
	// brandLogoFull	String	商品品牌logo全路径地址
	BrandLogoFull string `json:"brandLogoFull"`
	// schemeEndTime	Long	商品推广计划有效期预估截止时间：仅为预估时间，仅做参考；时间戳，单位：毫秒
	SchemeEndTime int64 `json:"schemeEndTime"`
	// sellTimeFrom	Long	商品售卖开始时间,时间戳，单位毫秒
	SellTimeFrom int64 `json:"sellTimeFrom"`
	// sellTimeTo	Long	商品售卖结束时间,时间戳, 单位毫秒
	SellTimeTo int64 `json:"sellTimeTo"`
	// weight	Integer	推广权重，用于确定推广该商品的优先级，权重值越大，优先级越高
	Weight int `json:"weight"`
	// storeInfo	StoreInfo	店铺信息
	StoreInfo struct {
		// storeId string		店铺id
		StoreID string `json:"storeId"`
		// storeName string		店铺名称
		StoreName string `json:"storeName"`
	} `json:"storeInfo"`
	// commentsInfo	GoodsCommentsInfo	商品评价信息
	CommentsInfo struct {
		// comments	Integer	商品评论数
		Comments int `json:"comments"`
		// goodCommentsShare	String	商品好评率:百分比，不带百分号
		GoodCommentsShare string `json:"goodCommentsShare"`
	} `json:"commentsInfo"`
	// storeServiceCapability	StoreServiceCapability	商品所属店铺服务能力评价
	StoreServiceCapability struct {
		// storeScore	string	店铺评分：保留两位小数
		StoreScore string `json:"storeScore"`
		// storeRankRate	string	店铺同品类排名比例：例如10表示前10%
		StoreRankRate string `json:"storeRankRate"`
	} `json:"storeServiceCapability"`
	// brandId	Long	商品所属档期（专场）id
	BrandID int64 `json:"brandId"`
	// schemeStartTime	Long	商品所属推广方案开始时间：时间戳，单位：毫秒
	SchemeStartTime int64 `json:"schemeStartTime"`
	// saleStockStatus	Integer	商品库存状态：1-已抢光，2-有库存，3-有机会,当列表查询库存或者查询商品详情时返回
	SaleStockStatus int `json:"saleStockStatus"`
	// status	Integer	商品状态：0-下架，1-上架
	Status int `json:"status"`
	// prepayInfo	PrepayInfo	商品预付信息
	PrepayInfo struct {
		// isPrepay	integer	是否预付商品:0-否，1-是
		IsPrepay int `json:"isPrepay"`
		// prepayPrice	string	预付到手价：元
		PrepayPrice string `json:"prepayPrice"`
		// firstAmount	string	预付首款金额：元
		FirstAmount string `json:"firstAmount"`
		// lastAmount	string	预付尾款金额：元
		LastAmount string `json:"lastAmount"`
		// prepayFavAmount	string	预付优惠金额: 元
		PrepayFavAmount string `json:"prepayFavAmount"`
		// deductionPrice	string	抵扣价格(首款+优惠金额)：元
		DeductionPrice string `json:"deductionPrice"`
		// prepayDiscount	string	预付折扣：(唯品价-优惠金额)/唯品价 保留两位小数的数字字符串
		PrepayDiscount string `json:"prepayDiscount"`
		// prepayFirstStartTime	long	首款支付开始时间:时间戳,单位毫秒
		PrepayFirstStartTime int64 `json:"prepayFirstStartTime"`
		// prepayFirstEndTime	long	首款支付结束时间:时间戳,单位毫秒
		PrepayFirstEndTime int64 `json:"prepayFirstEndTime"`
		// prepayLastStartTime	long	尾款支付开始时间:时间戳,单位毫秒
		PrepayLastStartTime int64 `json:"prepayLastStartTime"`
		// prepayLastEndTime	long	尾款支付结束时间:时间戳,单位毫秒
		PrepayLastEndTime int64 `json:"prepayLastEndTime"`
	} `json:"prepayInfo"`
	// joinedActivities	List<ActivityInfo>	商品参与活动信息:未参与活动集合为空
	JoinedActivities []struct {
		// actType	integer	活动类型:18-唯品快抢
		ActType int `json:"actType"`
		// actName	string	活动名称
		ActName string `json:"actName"`
		// beginTime	long	开始时间：时间戳，单位毫秒
		BeginTime int64 `json:"beginTime"`
		// endTime	long	结束时间：时间戳，单位毫秒
		EndTime int64 `json:"endTime"`
		// foreShowBeginTime	long	预热开始时间：时间戳，单位毫秒
		ForeShowBeginTime int64 `json:"foreShowBeginTime"`
	} `json:"joinedActivities"`
}
