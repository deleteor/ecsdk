package opentb

// TbkCouponGetRequest 阿里妈妈推广券信息查询
type TbkCouponGetRequest struct {
	// Me false	nfr%2BYTo2k1PX18gaNN%2BIPkIG2PadNYbBnwEsv6mRavWieOoOE3L9OdmbDSSyHbGxBAXjHpLKvZbL1320ML%2BCF5FRtW7N7yJ056Lgym4X01A%3D	带券ID与商品ID的加密串
	Me string `json:"me" form:"me"`
	// ItemID false	123	商品ID
	ItemID int `json:"item_id" form:"item_id"`
	// ActivityID false	sdfwe3eefsdf	券ID
	ActivityID string `json:"activity_id" form:"activity_id"`
}

// GetTbkCoupon 阿里妈妈推广券信息查询
func (c *Client) GetTbkCoupon(in *TbkCouponGetRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TbkCouponGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
