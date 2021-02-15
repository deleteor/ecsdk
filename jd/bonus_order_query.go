package jd

import (
	"encoding/json"
	"strings"
)

// BonusOrderParam 请求业务参数
type BonusOrderParam struct {
	OrderReq BonusOrderReq `json:"orderReq"`
}

// BonusOrderReq ...
type BonusOrderReq struct {
	// OptType 	是 1 时间类型 (1：下单时间，sortValue和pageSize组合使用； 2：更新时间，pageNo和pageSize组合使用)
	OptType int `json:"optType"`
	// StartTime  是 1541733000000 订单开始时间，时间戳（毫秒），起止时间限制10min内
	StartTime int64 `json:"startTime"`
	// EndTime  是 1541733600000 订单结束时间，时间戳（毫秒），起止时间限制10min内
	EndTime int64 `json:"endTime"`
	// PageNo  否 1 页码，默认值为1
	PageNo int `json:"pageNo"`
	// PageSize  是 100 每页数量，上限100
	PageSize int `json:"pageSize"`
	// SortValue  否 1541733600000,263989501198192640 时间类型按'下单'查询时，和pageSize组合使用。获取最后一条记录的sortValue，指定拉取条数（pageSize），以此方式查询数据。
	SortValue string `json:"sortValue"`
}

// JdUnionOpenOrderBonusQueryResponse ..
type JdUnionOpenOrderBonusQueryResponse struct {
	JdUnionOpenOrderBonusQueryResponse struct {
		Result string `json:"result"`
		Code   string `json:"code"`
	} `json:"jd_union_open_order_bonus_query_response"`
}

// BonusOrderResult 返回结果
type BonusOrderResult struct {
	Code      int               `json:"code"`
	Message   string            `json:"message"`
	HasMore   bool              `json:"hasMore"`
	RequestID string            `json:"requestId"`
	Data      []*BonusOrderResq `json:"data"`
}

// BonusOrderResq ...
type BonusOrderResq struct {
	// UnionID 	是	1000009196	联盟ID
	UnionID int64 `json:"unionId"`
	// BonusInvalidCode 	是	3	无效状态码，-1:无效、2:无效-拆单、3:无效-取消、4:无效-京东帮帮主订单、5:无效-账号异常、6:无效-赠品类目不返佣 等
	BonusInvalidCode string `json:"bonusInvalidCode"`
	// BonusInvalidText 	是	无效-取消	无效状态码对应的无效状态文案
	BonusInvalidText string `json:"bonusInvalidText"`
	// PayPrice 	是	0.00	实际支付金额
	PayPrice float64 `json:"payPrice"`
	// EstimateCosPrice 	是	1.00	预估计佣金额
	EstimateCosPrice float64 `json:"estimateCosPrice"`
	// EstimateFee 	是	0.04	预估佣金
	EstimateFee float64 `json:"estimateFee"`
	// ActualCosPrice 	是	0.00	实际计佣金额
	ActualCosPrice float64 `json:"actualCosPrice"`
	// ActualFee 	是	0.00	实际佣金
	ActualFee float64 `json:"actualFee"`
	// OrderTime 	是	1541733600000	下单时间、时间戳（毫秒）
	OrderTime int64 `json:"orderTime"`
	// FinishTime 	是	0	完成时间、时间戳（毫秒）
	FinishTime int64 `json:"finishTime"`
	// PositionID 	是	502047534	推广位ID
	PositionID int64 `json:"positionId"`
	// OrderID 	是	81823969001	订单号
	OrderID int64 `json:"orderId"`
	// BonusState 	是	0	奖励状态，0:无效、1:有效
	BonusState int `json:"bonusState"`
	// BonusText 	是	无效	奖励状态文案
	BonusText string `json:"bonusText"`
	// SkuName 	是	【京东拼购】拼团户外手套男骑行保暖女运动防滑加绒触屏手套 颜色款式随机发 S/M	商品名称
	SkuName string `json:"skuName"`
	// CommissionRate 	是	4.0	佣金比例，单位：%
	CommissionRate float64 `json:"commissionRate"`
	// SubUnionId 	是	LX_STORE_47534	子联盟ID
	SubUnionID string `json:"subUnionId"`
	// Pid 	是	3131_131_11	pid
	Pid string `json:"pid"`
	// Ext1 	是	无	推客生成推广链接时传入的扩展字段
	Ext1 string `json:"ext1"`
	// UnionAlias 	是	fhakf	母站长简称
	UnionAlias string `json:"unionAlias"`
	// SubSideRate 	是	95.00	分成比例，单位：%
	SubSideRate float64 `json:"subSideRate"`
	// SubsidyRate 	是	0.00	补贴比例，单位：%
	SubsidyRate float64 `json:"subsidyRate"`
	// FinalRate 	是	95.00	最终比例，单位：%
	FinalRate float64 `json:"finalRate"`
	// ActivityName 	是	11月京粉“拼购节”重点媒体奖励活动	活动名称
	ActivityName string `json:"activityName"`
	// ParentID 	是	81823969001	parentId
	ParentID int64 `json:"parentId"`
	// SkuId 	是	35847800771	skuId
	SkuID int64 `json:"skuId"`
	// EstimateBonusFee 	是	0	预估奖励金额
	EstimateBonusFee float64 `json:"estimateBonusFee"`
	// ActualBonusFee 	是	0	实际奖励金额
	ActualBonusFee float64 `json:"actualBonusFee"`
	// OrderState 	是	2	订单奖励状态，1:已完成、2:已付款、3:待付款
	OrderState int `json:"orderState"`
	// OrderText 	是	已付款	订单奖励状态
	OrderText string `json:"orderText"`
	// SortValue 	是	1541733600000,263989501198192640	排序值，按'下单时间'分页查询时使用
	SortValue string `json:"sortValue"`
}

// GetBonusOrders 获取奖励订单
func (J *Jdsdk) GetBonusOrders(ParamJsons string) (result *BonusOrderResult) {
	Method := "jd.union.open.order.bonus.query"
	J.SetSignJointUrlParam(Method, ParamJsons)
	var urls strings.Builder
	urls.WriteString(JD_HOST)
	urls.WriteString(J.SignAndUri)
	body, _ := HttpGet(urls.String())
	// fmt.Println(urls.String())
	response := &JdUnionOpenOrderBonusQueryResponse{}
	if len([]byte(body)) == 0 {
		return nil
	}
	e := json.Unmarshal([]byte(body), &response)
	if e != nil {
		panic(e)
	}
	result = &BonusOrderResult{}
	e = json.Unmarshal([]byte(response.JdUnionOpenOrderBonusQueryResponse.Result), result)
	if e != nil {
		panic(e)
	}
	return result
}
