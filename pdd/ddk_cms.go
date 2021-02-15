package pdd

import (
	"encoding/json"
)

// CMSPromURL .
type CMSPromURL struct {
	URL                      string `json:"url"`
	ShortURL                 string `json:"short_url"`
	MobileURL                string `json:"mobile_url"`
	MobileShortURL           string `json:"mobile_short_url"`
	WeAppWebViewURL          string `json:"we_app_web_view_url"`
	WeAppWebViewShortURL     string `json:"we_app_web_view_short_url"`
	MultiGroupURL            string `json:"multi_group_url"`              // 红包推广多人团链接
	MultiGroupShortURL       string `json:"multi_group_short_url"`        // 红包推广多人团短链接
	MultiGroupMobileURL      string `json:"multi_group_mobile_url"`       // 红包推广多人团移动链接
	MultiGroupMobileShortURL string `json:"multi_group_mobile_short_url"` // 红包推广多人团移动短链接
}

// ChannelType int
type ChannelType int

//
const (
	PackageMail19  ChannelType = 0 // 1.9 包邮
	TodayHotStyle  ChannelType = 1 // 今日爆款
	BrandClearance ChannelType = 2 // 品牌清仓
	PCTerminalMall ChannelType = 4 // PC端专属商城
	RaiseBabyCash  ChannelType = 5 // 养宝宝兑现金
)

// CMSPromURLGen .
func (d *DDK) CMSPromURLGen(pids []string, channelType ChannelType, generateWeappWebview, weAppWebViewShortURL, weAppWebViewURL bool, notMustparams ...Params) (res *CMSPromURL, err error) {
	res = new(CMSPromURL)
	params := NewParamsWithType(DDKCMSPromURLGenerate, notMustparams...)
	params.Set("p_id_list", TransformPids(pids))
	params.Set("channel_type", int(channelType))
	params.Set("generate_weapp_webview", generateWeappWebview)
	params.Set("we_app_web_view_url", weAppWebViewShortURL)
	params.Set("we_app_web_view_short_url", weAppWebViewURL)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "cms_promotion_url_generate_response", "url_list")
	err = json.Unmarshal(bytes, &res)
	return
}
