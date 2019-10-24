package opentb

// OpenuidGetBytradeRequest 请求参数
type OpenuidGetBytradeRequest struct {
	//TID true	123123	订单ID
	TID int `json:"tid" form:"tid"`
}

// GetOpenuidBytrade 通过订单获取对应买家的openUID
func (c *Client) GetOpenuidBytrade(in *OpenuidGetBytradeRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(OpenuidGetBytrade)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
