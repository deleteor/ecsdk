package opentb

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DoRequest 开始请求  传入参数：一个包含所有参数的map  返回参数：返回一个map
func (c *Client) DoRequest(params map[string]interface{}) (map[string]interface{}, error) {
	commonParams := c.SetCommonParams()
	requestParams := c.SetRequestParams(params, commonParams)

	// 调用签名
	if requestParams["sign_method"] == "md5" {
		sign := c.SignMD5(requestParams)
		requestParams["sign"] = sign
	}
	if requestParams["sign_method"] == "hmac" {
		sign := c.SignHMAC(requestParams)
		requestParams["sign"] = sign
	}

	values := url.Values{}
	for k, v := range requestParams {
		r := ""
		switch v.(type) {
		case int:
			r = strconv.Itoa(v.(int))
		case string:
			r = v.(string)
		}
		values.Set(k, r)
	}

	requestURL := ""
	if c.UseHTTPS != 0 {
		requestURL = httpsURL
	} else {
		requestURL = httpURL
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	requestClient := &http.Client{Transport: tr}

	requestBody, err := c.MakeRequestBody(requestParams)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", requestURL, requestBody)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := requestClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseStr := string(body)
	if strings.Contains(responseStr, "error") {

	}
	tbPwd := make(map[string]interface{})
	json.Unmarshal(body, &tbPwd)
	return tbPwd, nil
}
