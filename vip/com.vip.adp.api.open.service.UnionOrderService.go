package vip

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// UnionOrderServiceName 服务名
	UnionOrderServiceName string = "com.vip.adp.api.open.service.UnionOrderService"
	// UnionOrderVersion 版本号
	UnionOrderVersion string = "1.0.0"
	// UnionOrderMethodName 获取订单信息列表
	UnionOrderMethodName string = "orderList"
	// RefundOrderMethodName 获取订单信息列表
	RefundOrderMethodName string = "refundOrderList"
)

// OrderListQueryModel 获取订单信息列表 应用级请求参数
type OrderListQueryModel struct {
	// Status  否 订单状态:0-不合格，1-待定，2-已完结，该参数不设置默认代表全部状态
	Status uint `json:"status,omitempty"`
	// OrderTimeStart  否 订单时间起始 时间戳 单位毫秒
	OrderTimeStart int64 `json:"orderTimeStart,omitempty"`
	// OrderTimeEnd  否 订单时间结束 时间戳 单位毫秒
	OrderTimeEnd int64 `json:"orderTimeEnd,omitempty"`
	// Page  是 页码：从1开始
	Page int `json:"page,omitempty"`
	// PageSize  否 页面大小：默认20
	PageSize int `json:"pageSize,omitempty"`
	// RequestID  是 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	RequestID string `json:"requestId"`
	// UpdateTimeStart  否 更新时间-起始 时间戳 单位毫秒
	UpdateTimeStart int64 `json:"updateTimeStart,omitempty"`
	// UpdateTimeEnd  否 下单时间-结束 时间戳 单位毫秒
	UpdateTimeEnd int64 `json:"updateTimeEnd,omitempty"`
	// OrderSnList 否 订单号列表：当传入订单号列表时，订单时间和更新时间区间可不传入
	OrderSnList []string `json:"orderSnList,omitempty"`
}

// QueryOrderList 获取订单信息列表
func (c *Client) QueryOrderList(in *OrderListQueryModel) (*OrderListResponse, error) {

	c.SetMethod(UnionOrderMethodName)
	c.SetServiceName(UnionOrderServiceName)
	c.SetVersion(UnionOrderVersion)

	params := GetParams(in)
	// fmt.Println("params", params)
	mapParams := make(map[string]interface{})
	mapParams["queryModel"] = params
	result, err := c.DoRequest(mapParams)
	if err != nil {
		return nil, err
	}

	out := &UnionOrderServiceOrderListResponse{}
	outByte, _ := json.Marshal(result)
	json.Unmarshal(outByte, out)
	if out.ReturnCode != 0 {
		er := fmt.Sprintf("获取订单信息失败 code=%d\n", out.ReturnCode)
		return nil, errors.New(er)
	}
	return &out.Result, nil
}

// QueryRefundOrderList 获取订单信息列表
func (c *Client) QueryRefundOrderList(in *RefundOrderRequest) (*RefundOrderInfoList, error) {

	c.SetServiceName(UnionOrderServiceName)
	c.SetVersion(UnionOrderVersion)
	c.SetMethod(RefundOrderMethodName)

	params := GetParams(in)
	// fmt.Println("params", params)
	mapParams := make(map[string]interface{})
	mapParams["request"] = params
	result, err := c.DoRequest(mapParams)
	if err != nil {
		return nil, err
	}

	out := &UnionOrderServiceRefundOrderResponse{}
	outByte, _ := json.Marshal(result)
	json.Unmarshal(outByte, out)
	if out.ReturnCode != 0 {
		er := fmt.Sprintf("获取订单信息失败 code=%d\n", out.ReturnCode)
		return nil, errors.New(er)
	}
	return &out.Result, nil
}

// UnionOrderServiceOrderListResponse 订单响应
type UnionOrderServiceOrderListResponse struct {
	ReturnCode int               `json:"returnCode"`
	Result     OrderListResponse `json:"result"`
}

// OrderListResponse 订单响应参数
type OrderListResponse struct {
	// orderInfoList List<OrderInfo> 否 业绩查询响应结果
	OrderInfoList []OrderInfo `json:"orderInfoList"`
	// total Integer 否 业绩总条数
	Total int `json:"total"`
	// page Integer 否 当前页码
	Page int `json:"page"`
	// pageSize Integer 否 页面大小
	PageSize int `json:"pageSize"`
}

// OrderInfo 订单详情
type OrderInfo struct {
	// orderSn	String	否	订单号
	OrderSn string `json:"orderSn"`
	// status	Short	否	订单状态:0-不合格，1-待定，2-已完结
	Status uint `json:"status"`
	// newCustomer	Short	否	新老客：0-待定，1-新客，2-老客
	NewCustomer uint `json:"newCustomer"`
	// channelTag	String	否	渠道标识
	ChannelTag string `json:"channelTag"`
	// orderTime	Long	否	下单时间 时间戳 单位毫秒
	OrderTime int64 `json:"orderTime"`
	// signTime	Long	否	签收时间 时间戳 单位毫秒
	SignTime int64 `json:"signTime"`
	// settledTime	Long	否	结算时间 时间戳 单位毫秒
	SettledTime int64 `json:"settledTime"`
	// detailList	List<OrderDetailInfo>	否	商品明细
	DetailList []OrderDetailInfo `json:"detailList"`
	// lastUpdateTime	Long	否	订单上次更新时间 时间戳 单位毫秒
	LastUpdateTime int64 `json:"lastUpdateTime"`
	// settled	Short	否	订单结算状态 0-未结算,1-已结算
	Settled uint `json:"settled"`
	// selfBuy	Integer	否	是否自推自买 0-否，1-是
	SelfBuy int `json:"selfBuy"`
	// orderSubStatusName	String	否	订单子状态：流转状态-支持状态：（已下单、已付款、已签收、待结算、已结算、已失效）
	OrderSubStatusName string `json:"orderSubStatusName"`
	// commission	String	否	商品总佣金:单位元
	Commission string `json:"commission"`
	// afterSaleChangeCommission	String	否	售后订单佣金变动：仅在订单完结之后发生售后行为时返回
	AfterSaleChangeCommission string `json:"afterSaleChangeCommission"`
	// afterSaleChangeGoodsCount	Integer	否	售后订单总商品数量变动：仅在订单完结之后发生售后行为时返回
	AfterSaleChangeGoodsCount int `json:"afterSaleChangeGoodsCount"`
	// commissionEnterTime	Long	否	入账时间，时间戳，单位毫秒
	CommissionEnterTime int64 `json:"commissionEnterTime"`
	// orderSource	String	否	订单来源
	OrderSource string `json:"orderSource"`
	// pid	String	否	推广PID:目前等同于channelTag
	Pid string `json:"pid"`
	// isPrepay	Integer	否	是否预付订单:0-否，1-是
	IsPrepay int `json:"isPrepay"`
}

// OrderDetailInfo 订单商品明细
type OrderDetailInfo struct {
	// goodsId	String	否	商品id
	GoodsID string `json:"goodsId"`
	// goodsName	String	否	商品名称
	GoodsName string `json:"goodsName"`
	// goodsThumb	String	否	商品缩略图
	GoodsThumb string `json:"goodsThumb"`
	// goodsCount	Integer	否	商品数量
	GoodsCount int `json:"goodsCount"`
	// commissionTotalCost	String	否	商品计佣金额
	CommissionTotalCost string `json:"commissionTotalCost"`
	// commissionRate	String	否	商品佣金比例
	CommissionRate string `json:"commissionRate"`
	// commission	String	否	商品佣金金额
	Commission string `json:"commission"`
	// commCode	String	否	佣金编码：对应商品二级分类
	CommCode string `json:"commCode"`
	// commName	String	否	佣金方案名称
	CommName string `json:"commName"`
	// orderSource	String	否	订单来源
	OrderSource string `json:"orderSource"`
	// afterSaleInfo	List<AfterSaleDetailInfo>	否	商品售后信息
	AfterSaleInfo []AfterSaleDetailInfo `json:"afterSaleInfo"`
	// sizeId	String	否	商品尺码：2019.01.01之后可用
	SizeID string `json:"sizeId"`
	// status	Short	否	商品状态：0-不合格，1-待定，2-已完结
	Status uint `json:"status"`
}

// AfterSaleDetailInfo 商品售后信息
type AfterSaleDetailInfo struct {
	// afterSaleChangedCommission	String	否	商品佣金售后变动:仅在订单完结之后发生售后时返回，无售后时为空
	AfterSaleChangedCommission string `json:"afterSaleChangedCommission"`
	// afterSaleChangedGoodsCount	Integer	否	商品数量售后变动:仅在订单完结之后发生售后时返回，无售后时为空
	AfterSaleChangedGoodsCount int `json:"afterSaleChangedGoodsCount"`
	// afterSaleSn	String	否	商品售后单号，无售后时为空
	AfterSaleSn string `json:"afterSaleSn"`
	// afterSaleStatus	Integer	否	商品售后状态：1-售后中，2-售后完成，3-售后取消，无售后时为空
	AfterSaleStatus int `json:"afterSaleStatus"`
	// afterSaleType	Integer	否	售后类型：1-退货，2-换货,无售后时为空
	AfterSaleType int `json:"afterSaleType"`
	// afterSaleFinishTime	Long	否	售后完成时间，时间戳，单位：毫秒，无售后时为空
	AfterSaleFinishTime int64 `json:"afterSaleFinishTime"`
}

// RefundOrderRequest 联盟维权订单请求参数
type RefundOrderRequest struct {
	// searchType Integer 查询类型:0-维权发起时间，1-维权完成时间（佣金扣除入账时间），2-佣金扣除结算时间
	SearchType int `json:"searchType"`
	// searchTimeStart Long 目标时间起始：时间戳，单位毫秒; 当searchType=0时，为维权发起时间， 当searchType=1时，为维权完成时间（佣金扣除入账时间）， 当searchType=2时，为佣金扣除结算时间
	SearchTimeStart int64 `json:"searchTimeStart"`
	// searchTimeEnd Long 目标时间结束：时间戳，单位毫秒; 当searchType=0时，为维权发起时间， 当searchType=1时，为维权完成时间（佣金扣除入账时间）， 当searchType=2时，为佣金扣除结算时间
	SearchTimeEnd int64 `json:"searchTimeEnd"`
	// orderSns List<String> 目标订单号集合:当指定订单号集合时，目标时间区间可以不传,否则必须指定目标时间区间
	OrderSns []string `json:"orderSns"`
	// page Integer 当前页
	Page int `json:"page"`
	// pageSize Integer 分页大小：默认20，最大100
	PageSize int `json:"pageSize"`
	// requestId String 请求id：调用方自行定义，用于追踪请求，单次请求唯一，建议使用UUID
	RequestID string `json:"requestId"`
}

// UnionOrderServiceRefundOrderResponse 订单响应
type UnionOrderServiceRefundOrderResponse struct {
	ReturnCode int                 `json:"returnCode"`
	Result     RefundOrderInfoList `json:"result"`
}

// RefundOrderInfoList 联盟维权订单信息
type RefundOrderInfoList struct {
	// refundOrderInfoList List<RefundOrderInfo> 联盟维权订单信息
	RefundOrderInfoList []RefundOrderInfo `json:"refundOrderInfoList"`
	// total Integer 总记录数
	Total int `json:"total"`
	// page Integer 当前页
	Page int `json:"page"`
	// pageSize Integer 分页大小
	PageSize int `json:"pageSize"`
}

// RefundOrderInfo 联盟维权订单信息
type RefundOrderInfo struct {
	// orderSn	String	订单号
	OrderSn string `json:"orderSn"`
	// applyTime	Long	申请时间:时间戳，单位毫秒
	ApplyTime int64 `json:"applyTime"`
	// refundType	Integer	维权类型：1-退货，2-换货
	RefundType int `json:"refundType"`
	// commissionTotalCost	String	订单计佣金额:单位元
	CommissionTotalCost string `json:"commissionTotalCost"`
	// commission	String	订单佣金：单位元
	Commission string `json:"commission"`
	// goodsCount	Integer	商品数量
	GoodsCount int `json:"goodsCount"`
	// commissionEnterTxn	String	维权返还佣金的入账流水号
	CommissionEnterTxn string `json:"commissionEnterTxn"`
	// commissionEnterTime	Long	维权返还佣金的入账时间: 时间戳，单位毫秒
	CommissionEnterTime int64 `json:"commissionEnterTime"`
	// commissionSettledTime	Long	维权返回佣金的结算时间：时间戳，单位毫秒
	CommissionSettledTime int64 `json:"commissionSettledTime"`
	// refundOrderDetails	List<RefundOrderDetail>	维权订单明细
	RefundOrderDetails []RefundOrderDetail `json:"refundOrderDetails"`
	// userId	Long	订单归属人id
	UserID int64 `json:"userId"`
	// orderTime	Long	订单时间:时间戳，单位毫秒
	OrderTime int64 `json:"orderTime"`
	// originCommEnterTime	Long	订单入账时间：发起维权之前入账时间，时间戳，单位毫秒
	OriginCommEnterTime int64 `json:"originCommEnterTime"`
	// originCommEnterTxn	String	订单入账流水：发起维权之前入账流水号
	OriginCommEnterTxn string `json:"originCommEnterTxn"`
	// settled	Integer	售后订单结算状态：0-未结算，1-已结算
	Settled int `json:"settled"`
	// afterSaleSn	String	售后单号
	AfterSaleSn string `json:"afterSaleSn"`
	// afterSaleStatus	Short	售后单状态:1-售后中，2-售后完成，3-售后取消
	AfterSaleStatus uint `json:"afterSaleStatus"`
}

// RefundOrderDetail 维权订单明细
type RefundOrderDetail struct {
	// goodsId	String	商品id
	GoodsID string `json:"goodsId"`
	// sizeId	String	商品尺码id
	SizeID string `json:"sizeId"`
	// goodsPrice	String	商品价格：元
	GoodsPrice string `json:"goodsPrice"`
	// goodsCount	Integer	维权商品数量
	GoodsCount int `json:"goodsCount"`
	// commission	String	维权商品佣金
	Commission string `json:"commission"`
	// totalCost	String	维权商品计佣金额
	TotalCost string `json:"totalCost"`
}
