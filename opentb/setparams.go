package opentb

import (
	"fmt"
	"io"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Client 请求体
type Client struct {
	AppKey     string
	AppSecret  string
	UseHTTPS   bool
	Method     string
	SignMethod string
	Format     string
	NumIids    string
}

const (
	// httpURL is for HTTP REST API URL.
	httpURL string = "http://gw.api.taobao.com/router/rest"
	// httpsURL is for HTTPS REST API URL.
	httpsURL string = "https://eco.taobao.com/router/rest"
)

// SetCommonParams 设置通用的Map
func (c *Client) SetCommonParams() map[string]interface{} {
	params := make(map[string]interface{})
	t := time.Now()
	params["format"] = c.Format
	params["v"] = "2.0"
	params["sign_method"] = c.SignMethod
	params["app_key"] = c.AppKey
	params["timestamp"] = fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	return params
}

// SetRequestParams 生成全部参数
func (c *Client) SetRequestParams(requestParams, commonParams map[string]interface{}) map[string]interface{} {
	for k := range commonParams {
		requestParams[k] = commonParams[k]
		// fmt.Printf("---%s", requestParams)
	}
	fmt.Printf("%s", requestParams)
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
		case string:
			r = v.(string)
		}
		values.Set(k, r)
	}
	return strings.NewReader(values.Encode()), nil
}
