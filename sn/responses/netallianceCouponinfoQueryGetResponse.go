package responses

type NetallianceCouponinfoQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			QueryCouponinfo QueryCouponinfo `json:"queryCouponinfo"`
		} `json:"sn_body"`
		ErrorResponse
	} `json:"sn_responseContent"`
}

type QueryCouponinfo struct {
	ActivityId            string `json:"activityId"`            //券活动ID
	CouponValue           string `json:"couponValue"`           //券面额
	CouponCount           string `json:"couponCount"`           //券总张数
	CouponRemainingAmount string `json:"couponRemainingAmount"` //券剩余数量
	PreferentialDistinct  string `json:"preferentialDistinct"`  // 奖励优惠说明
	CouponStartTime       string `json:"CouponStartTime"`       //券的使用开始时间
	CouponEndTime         string `json:"couponEndTime"`         //券的使用结束时间
	StartTime             string `json:"startTime"`             //券的领取开始时间
	EndTime               string `json:"endTime"`               //券的领取结束时间
	CouponPlatform        string `json:"couponPlatform"`        //自营    券平台
	IsValidCoupon         string `json:"isValidCoupon"`         //0代表有效1代表无效
}
