package vip

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// DoRequest 执行请求
func (c *Client) DoRequest(params map[string]interface{}) (map[string]interface{}, error) {
	commonParams := c.SetCommonParams()
	strParams, _ := json.Marshal(params)
	// 调用签名
	c.Sign = c.HMACMD5(commonParams, string(strParams))
	commonParams["sign"] = c.Sign

	values := url.Values{}
	for k, v := range commonParams {
		r := ""
		switch v.(type) {
		case int:
			r = strconv.Itoa(v.(int))
		case string:
			r = v.(string)
		}
		values.Set(k, r)
	}

	// 拼接url
	requestURL := ""
	if c.UseHTTPS != false {
		requestURL = httpsURL + "?" + values.Encode()
	} else {
		requestURL = httpURL + "?" + values.Encode()
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	requestClient := &http.Client{Transport: tr}

	requestBody := marshal(params)
	req, err := http.NewRequest("POST", requestURL, strings.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

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
	info := make(map[string]interface{})
	json.Unmarshal(body, &info)
	return info, nil
}

func marshal(obj interface{}) string {
	var bytes, err = json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(bytes)
}
