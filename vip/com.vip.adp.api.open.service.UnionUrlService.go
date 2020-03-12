package vip

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// UnionURLServiceName 服务名
	UnionURLServiceName string = "com.vip.adp.api.open.service.UnionUrlService"
	// UnionURLVersion 版本号
	UnionURLVersion string = "1.0.0"
	// UnionURLMethodNameGenByVIPUrl 获取订单信息列表
	UnionURLMethodNameGenByVIPUrl string = "genByVIPUrl"
)

// URLListRequest 获取订单信息列表 应用级请求参数
type URLListRequest struct {
	// urlList	List<String>	商品链接列表
	URLList []string `json:"urlList"`
	// chanTag	String	渠道标识
	ChanTag string `json:"chanTag	String"`
	// requestId	String	请求id：UUID
	RequestID string `json:"requestId	String"`
}

// genByVIPUrl 获取订单信息列表
func (c *Client) genByVIPUrl(in *URLListRequest) ([]*URLInfo, error) {

	c.SetMethod(UnionURLServiceName)
	c.SetVersion(UnionURLVersion)
	c.SetServiceName(UnionURLMethodNameGenByVIPUrl)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}

	fmt.Println("===", result)
	out := &URLGenResponseResponse{}
	outByte, _ := json.Marshal(result)
	json.Unmarshal(outByte, out)
	if out.ReturnCode != 0 {
		er := fmt.Sprintf("获取订单信息失败 code=%d\n", out.ReturnCode)
		return nil, errors.New(er)
	}
	return out.Result.URLInfoList, nil
}

// URLGenResponseResponse 订单响应
type URLGenResponseResponse struct {
	ReturnCode int            `json:"returnCode"`
	Result     URLGenResponse `json:"result"`
}

// URLGenResponse ..
type URLGenResponse struct {
	URLInfoList []*URLInfo `json:"urlInfoList"`
}

// URLInfo 链接数据
type URLInfo struct {
	// source string 链接生成的数据源： 如果targetType为GOODSID，该值商品id， 如果targetType为URL，该值为唯品会url， 如果targetType为VIP_SMALL和GENERAL_SAMLL时，该值为空（此类型VOP暂未开放）
	Source string `json:"source"`
	// uRL string CPS短链接
	URL string `json:"url"`
	// longURL string CPS长连接
	LongURL string `json:"longUrl"`
	// ulURL string CPS通用连接
	UlURL string `json:"ulUrl"`
	// deeplinkURL string CPS Deeplink链接
	DeeplinkURL string `json:"deeplinkUrl"`
	// traFrom string 小程序CPS参数，如果targetType为VIP_SMALL和GENERAL_SAMLL时，该参数将被赋值
	TraFrom string `json:"traFrom"`
	// noEvokeURL string CPS短链接：不唤起快应用
	NoEvokeURL string `json:"noEvokeUrl"`
	// noEvokeLongURL string CPS长链接：不唤起快应用
	NoEvokeLongURL string `json:"noEvokeLongUrl"`
	// vipWxURL string 唯品会小程序链接：仅在根据商品id获取时返回
	VipWxURL string `json:"vipWxUrl"`
	// vipWxCode string 唯品会小程序码：仅在根据商品id获取时返回
	VipWxCode string `json:"vipWxCode"`
	// vipQuickAppURL string 唯品会快应用链接
	VipQuickAppURL string `json:"vipQuickAppUrl"`
}
