package vip

import (
	"encoding/json"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Client 请求体
type Client struct {
	Service   string
	Method    string
	Version   string
	Timestamp string
	Format    string
	AppKey    string
	AppSecret string
	UseHTTPS  bool
	Sign      string
}

const (
	// httpURL is for HTTP REST API URL.
	httpURL string = "http://sandbox.vipapis.com"
	// httpsURL is for HTTPS REST API URL.
	httpsURL string = "https://gw.vipapis.com"
)

// NewClient ..
func NewClient(AppKey, SecretKey string) *Client {
	return &Client{
		AppKey:    AppKey,
		AppSecret: SecretKey,
		UseHTTPS:  true,
		Format:    "JSON",
		Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
		// Timestamp: "1583834213",
	}
}

// SetAppKey 设置新配置
func (c *Client) SetAppKey(AppKey, AppSecret string) {
	c.AppKey, c.AppSecret = AppKey, AppSecret
}

// SetHTTPS 设置新配置
func (c *Client) SetHTTPS(f bool) {
	c.UseHTTPS = f
}

// SetServiceName 设置服务名
func (c *Client) SetServiceName(m string) {
	c.Service = m
}

// SetMethod 设置方法名
func (c *Client) SetMethod(m string) {
	c.Method = m
}

// SetVersion 设置方法名
func (c *Client) SetVersion(m string) {
	c.Version = m
}

// SetCommonParams 系统级参数
func (c *Client) SetCommonParams() map[string]interface{} {
	params := make(map[string]interface{})
	params["service"] = c.Service
	params["method"] = c.Method
	params["version"] = c.Version
	params["timestamp"] = c.Timestamp
	params["format"] = c.Format
	params["appKey"] = c.AppKey
	return params
}

// SortParamsToStr 给参数排序，并返回数值
func (c *Client) SortParamsToStr(params map[string]interface{}) string {
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str := ""
	for _, k := range keys {
		r := ""
		switch params[k].(type) {
		case int:
			r = strconv.Itoa(params[k].(int))
		case float64:
			r = strconv.FormatFloat(params[k].(float64), 'E', -1, 64)
		case string:
			r = params[k].(string)
		}
		str += k + r
	}
	return str
}

// MakeRequestBody 创建请求体
func (c *Client) MakeRequestBody(params map[string]interface{}) (io.Reader, error) {
	values := url.Values{}
	for k, v := range params {
		r := ""
		switch v.(type) {
		case int:
			r = strconv.Itoa(v.(int))
		case float64:
			r = strconv.FormatFloat(v.(float64), 'E', -1, 64)
		case string:
			r = v.(string)
		}
		values.Set(k, r)
	}
	return strings.NewReader(values.Encode()), nil
}

// GetParams 获取请求params 去除0值参数
func GetParams(in interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	j, _ := json.Marshal(in)
	json.Unmarshal(j, &params)
	// for k, par := range params {
	// 	switch par.(type) {
	// 	case string:
	// 		if par == "" {
	// 			delete(params, k)
	// 		}
	// 	case float64:
	// 		if par == 0.00 {
	// 			delete(params, k)
	// 		}
	// 	}
	// }
	return params
}
