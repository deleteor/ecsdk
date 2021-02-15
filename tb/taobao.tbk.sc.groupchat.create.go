package tb

// TbkScGroupchatCreateRequest 淘宝客-服务商-手淘群创建
type TbkScGroupchatCreateRequest struct {
	// Title	true	淘客推广群	群组名称
	Title string `json:"title" form:"title"`
	// SubGroupNum	false	10	目前一个淘客仅开放1个可以容纳5000人的大群，大群下面有10个小群，每个小群可以容纳500人；当小群加满之后会自动创建另一个新的小群，上限是10个小群
	SubGroupNum int `json:"sub_group_num" form:"sub_group_num"`
}

// CreateTbkScGroupchat 淘宝客-服务商-手淘群创建
func (c *Client) CreateTbkScGroupchat(in *TbkScGroupchatCreateRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScGroupchatCreate)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
