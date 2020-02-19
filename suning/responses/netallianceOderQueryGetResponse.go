package responses

type NetallianceOderQueryGetResponse struct {
	ResponseContent struct {
		ResponseBody struct {
			QueryOrder []OrderList `json:"queryOrder"`
		} `json:"sn_body"`
		ResponseHead struct {
			TotalSize     int64  `json:"totalSize"`
			PageTotal     int64  `json:"pageTotal"`
			PageNo        int64  `json:"pageNo"`
			ReturnMessage string `json:"returnMessage"`
		} `json:"sn_head"`
		ErrorResponse
	} `json:"sn_responseContent"`
}

type OrderList struct {
	OrderCode   string        `json:"orderCode"`
	OrderDetail []OrderDetail `json:"orderDetail"`
}

type OrderDetail struct {
	PayTime                   string `json:"payTime"`                   //支付时间，格式：yyyy-MM-dd HH:mm:ss
	OrderSubmitTime           string `json:"orderSubmitTime"`           //下单时间，格式：yyyy-MM-dd HH:mm:ss
	OrderLineNumber           string `json:"orderLineNumber"`           //订单行项目号
	OrderLineStatusDesc       string `json:"orderLineStatusDesc"`       //订单行项目状态
	OrderLineStatusChangeTime string `json:"orderLineStatusChangeTime"` //行项目状态更新时间,格式：yyyy-MM-dd HH:mm:ss
	OrderLineOrigin           string `json:"orderLineOrigin"`           //订单行来源（PC端、无线端）
	ProductName               string `json:"productName"`               //商品名称
	SaleNum                   string `json:"saleNum"`                   //商品数量
	PayAmount                 string `json:"payAmount"`                 //实付金额
	OrderLineFlag             string `json:"orderLineFlag"`             //订单行标记
	ChildAccountId            string `json:"childAccountId"`            //子推广账号ID(对应sub_user)
	SellName                  string `json:"sellName"`                  //商户名称
	SellerCode                string `json:"sellerCode"`                //商户编码
	GoodsNum                  string `json:"goodsNum"`                  //商品编码
	CommissionRatio           string `json:"commissionRatio"`           //佣金比例
	PrePayCommission          string `json:"prePayCommission"`          //预估佣金
	ProductFirstCatalog       string `json:"productFirstCatalog"`       //一级目录
	ProductSecondCatalog      string `json:"productSecondCatalog"`      //二级目录
	ProductThirdCatalog       string `json:"productThirdCatalog"`       //三级目录
	OrderType                 string `json:"orderType"`                 //商品归属
	PositionId                string `json:"positionId"`                //推广位ID
	GoodsGroupCatalog         string `json:"goodsGroupCatalog"`         //商品组目录编码
	SaleType                  string `json:"saleType"`                  //推广类型-链接推广
	PictureUrl                string `json:"pictureUrl"`                //图片
}
