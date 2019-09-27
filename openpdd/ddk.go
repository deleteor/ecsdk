package openpdd

import (
	"encoding/json"
)

type DDK struct {
	Context *Context
}

func NewDDK(cfg *Config) *DDK {
	return &DDK{Context: NewContext(cfg)}
}

func newDDKWithContext(ctx *Context) *DDK {
	return &DDK{Context: ctx}
}

type RedPackageURL struct {
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

func (d *DDK) RPPromURLGenerate(pids []string, generateWeappWebview bool, notMustparams ...Params) (res *RedPackageURL, err error) {
	res = new(RedPackageURL)
	params := NewParamsWithType(DDKRPPromURLGenerate, notMustparams...)
	params.Set("p_id_list", TransformPids(pids))
	params.Set("generate_weapp_webview", generateWeappWebview)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "rp_promotion_url_generate_response", "url_list")
	err = json.Unmarshal(bytes, &res)
	return
}

type LotteryURL struct {
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

// LotteryURLGen lottery custom_parameters
func (d *DDK) LotteryURLGen(pids []string, generateWeappWebview bool, notMustparams ...Params) (res *LotteryURL, err error) {
	res = new(LotteryURL)
	params := NewParamsWithType(DDKLotteryURLGen, notMustparams...)
	params.Set("pid_list", TransformPids(pids))
	params.Set("generate_weapp_webview", generateWeappWebview)

	r, err := Call(d.Context, params)
	if err != nil {
		return
	}
	bytes, err := GetResponseArrayIndexBytes(r, 0, "lottery_url_response", "url_list")
	err = json.Unmarshal(bytes, &res)
	return
}
