package opentb

import (
	"fmt"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"git.hp6h2.cn/gobin/commons/gtype/boolx"
)

// Client 请求体
type Client struct {
	AppKey     string
	AppSecret  string
	UseHTTPS   boolx.Bool
	Method     TaoBaoURI
	SignMethod string
	Format     string
	Session    string
}

const (
	// httpURL is for HTTP REST API URL.
	httpURL string = "http://gw.api.taobao.com/router/rest"
	// httpsURL is for HTTPS REST API URL.
	httpsURL string = "https://eco.taobao.com/router/rest"
)

// NewClient ..
func NewClient(AppKey, SecretKey string) *Client {
	return &Client{
		AppKey:     AppKey,
		AppSecret:  SecretKey,
		UseHTTPS:   1,
		SignMethod: "md5",
		Format:     "json",
	}
}

// SetAppKey 设置新配置
func (c *Client) SetAppKey(AppKey string, SecretKey string) {
	c.AppKey = AppKey
	c.AppSecret = SecretKey
}

// SetHTTPS 设置新配置
func (c *Client) SetHTTPS(f boolx.Bool) {
	c.UseHTTPS = f
}

// SetMethod 设置淘宝请求接口
func (c *Client) SetMethod(m TaoBaoURI) {
	c.Method = m
}

// SetSessionKey 设置淘宝请求接口
func (c *Client) SetSessionKey(session string) {
	c.Session = session
}

// SetCommonParams 设置通用的Map
func (c *Client) SetCommonParams() map[string]interface{} {
	params := make(map[string]interface{})
	t := time.Now()
	params["format"] = c.Format
	params["v"] = "2.0"
	params["method"] = c.Method
	params["sign_method"] = c.SignMethod
	params["app_key"] = c.AppKey
	params["timestamp"] = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	params["session"] = c.Session
	return params
}

// SetRequestParams 生成全部参数
func (c *Client) SetRequestParams(requestParams, commonParams map[string]interface{}) map[string]interface{} {
	for k := range commonParams {
		requestParams[k] = commonParams[k]
	}
	return requestParams
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
		case TaoBaoURI:
			r = string(params[k].(TaoBaoURI))
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
		case TaoBaoURI:
			r = string(v.(TaoBaoURI))
		}
		values.Set(k, r)
	}
	return strings.NewReader(values.Encode()), nil
}
