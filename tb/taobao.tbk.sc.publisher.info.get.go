package tb

// TbkScPublisherInfoGetRequest 淘宝客-公用-私域用户备案信息查询
type TbkScPublisherInfoGetRequest struct {
	// InfoType true	1	类型，必选 1:渠道信息；2:会员信息
	InfoType int `json:"info_type" form:"info_type"`
	// RelationId false	2323	渠道独占 - 渠道关系ID
	RelationID int `json:"relation_id" form:"relation_id"`
	// PageNo false	1	第几页
	PageNo int `json:"page_no" form:"page_no"`
	// PageSize false	10	每页大小
	PageSize int `json:"page_size" form:"page_size"`
	// RelationApp true	common	备案的场景：common（通用备案），etao（一淘备案），minietao（一淘小程序备案），offlineShop（线下门店备案），offlinePerson（线下个人备案）。如不填默认common。查询会员信息只需填写common即可
	RelationApp string `json:"relation_app" form:"relation_app"`
	// SpecialID false	1212	会员独占 - 会员运营ID
	SpecialID string `json:"special_id" form:"special_id"`
}

// GetTbkScPublisherInfo 淘宝客-公用-私域用户备案信息查询
func (c *Client) GetTbkScPublisherInfo(in *TbkScPublisherInfoGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkScPublisherInfoGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
