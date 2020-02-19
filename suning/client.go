package suning

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"gylm-example/thirdparty/suning/constants"
	"gylm-example/thirdparty/suning/requests"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Params map[string]interface{}

var (
	signMethod     = constants.SIGN_METHOD_MD5
	responseFormat = constants.FORMAT_JSON
	version        = constants.VERSION
	userAgent      = constants.USER_AGENT
	sdkVersion     = constants.SDK_VERSION
	serverUrl      = "https://open.suning.com/api/http/sopRequest"
)

// TopClient TOP默认客户端
type TopClient struct {
	appKey    string
	appSecret string
}

func verifyRequired(m map[string]string) error {
	for k, v := range m {
		if v == "" {
			return errors.New(k + " is required")
		}
	}
	return nil
}

// NewTopClient 创建 TopClient 的新实例
func NewTopClient(appKey, appSecret string) (*TopClient, error) {

	err := verifyRequired(map[string]string{
		"appKey":    appKey,
		"appSecret": appSecret,
	})

	if err != nil {
		return nil, err
	}

	return &TopClient{
		appKey, appSecret,
	}, nil
}

// Do 执行请求
func (c TopClient) Do(r requests.ITopRequest, respModel interface{}) error {

	r.Validate() //验证参数有效性

	form := r.GetForm()

	// 组装系统参数
	SysParams := make(Params)
	SysParams["secret_key"] = c.appSecret
	SysParams["method"] = r.APIName()
	date := time.Now().Format(constants.TIMESTAMP_FORMAT)
	SysParams["date"] = date
	SysParams["app_key"] = c.appKey
	SysParams["api_version"] = constants.VERSION

	// 获取业务参数
	paramsArray := make(map[string]interface{},1)
	paramsArray2 := make(map[string]interface{},1)
	paramsArray3 := make(map[string]interface{},1)
	paramsArray3[r.BizName()] = form
	paramsArray2["sn_body"] = paramsArray3
	paramsArray["sn_request"] = paramsArray2
	body := marshal(paramsArray)
	//fmt.Println(body)
	post_field := base64.URLEncoding.EncodeToString([]byte(body))
	SysParams["post_field"] = post_field
	//fmt.Println(SysParams)
	signString := sign(SysParams)
	requestURL := makeRequestURL(serverUrl, r.APIName())

	req, err := http.NewRequest("POST", requestURL,strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("AppMethod", r.APIName())
	req.Header.Set("AppRequestTime", date)
	req.Header.Set("Format", constants.FORMAT_JSON)
	req.Header.Set("signInfo", signString)
	req.Header.Set("AppKey", c.appKey)
	req.Header.Set("AppKey", c.appKey)
	req.Header.Set("VersionNo", constants.VERSION)
	req.Header.Set("User-Agent", constants.USER_AGENT)
	req.Header.Set("Sdk-Version", constants.SDK_VERSION)
	if err != nil {
		panic(err)
	}
	if header := r.GetHeader(); header != nil {
		req.Header = header
	}
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return parseResponseBody(resp, &respModel)
}

// parseResponseBody 解析响应结果的正文内容
func parseResponseBody(resp *http.Response, respModel interface{}) error {
	buf, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(buf))
	if err != nil {
		return err
	}

	switch responseFormat {
	case constants.FORMAT_JSON:
		return json.Unmarshal(buf, &respModel)
	case constants.FORMAT_XML:
		return  fmt.Errorf("unimplemented XML parser.")
		//panic("unimplemented XML parser.")
	default:
		return  fmt.Errorf("unimplemented " + responseFormat + " parser.")
		//panic("unimplemented " + responseFormat + " parser.")
	}
}

// makeRequestURL 构造请求地址
func makeRequestURL(serverUrl string, apiName string) string {
	return fmt.Sprintf("%v/%v", serverUrl, apiName)
}

// signForm 签名请求参数

func sign(sysParams Params) string {
	sortKeys := [6]string{"secret_key","method","date","app_key","api_version","post_field"}
	singStr := ""
	for _,key := range sortKeys {
		if val := getString(sysParams[key]); val != "" {
			singStr +=  val
		}
	}
	return createSign(singStr)
}

func createSign(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return string(hex.EncodeToString(cipherStr))
}

func getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
	return ""
}

func marshal(obj interface{}) string {
	var bytes, err = json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}
