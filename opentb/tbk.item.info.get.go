package opentb

import (
	"encoding/json"
)

// TbkItemInfoGetRequest 请求参数
type TbkItemInfoGetRequest struct {
	NumIids  string `json:"num_iids" form:"num_iids"`
	Platform int    `json:"platform" form:"platform"`
	IP       string `json:"ip" form:"ip"`
}

// GetTbkItemInfo 淘宝客-公用-淘宝客商品详情查询(简版)
func (c *Client) GetTbkItemInfo(in *TbkItemInfoGetRequest, cfg *Config) (map[string]interface{}, error) {
	// c.Method = TbkItemInfoGet
	c.SetMethod(TbkItemInfoGet)
	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetParams 获取请求params 去除0值参数
func GetParams(in interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	j, _ := json.Marshal(in)
	json.Unmarshal(j, &params)
	for k, par := range params {
		switch par.(type) {
		case string:
			if par == "" {
				delete(params, k)
			}
		case float64:
			if par == 0.00 {
				delete(params, k)
			}
		}
	}
	return params
}
