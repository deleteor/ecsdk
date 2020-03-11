package vip

import (
	"testing"
)

func TestVip(t *testing.T) {

	cfg := &Config{
		AppKey:    "6e96161d",
		SecretKey: "59F170C722E1BADDCF574E056622B2B0",
	}
	c := NewClient(cfg.AppKey, cfg.SecretKey)

	in := &OrderListQueryModel{
		Status:         0,
		OrderTimeStart: 1580547124,
		OrderTimeEnd:   1583052724,
		Page:           1,
		PageSize:       20,
		RequestID:      "1",
	}
	res, err := c.QueryOrderList(in)
	if err != nil {
		t.Errorf("%+v", err)
		return
	}
	t.Logf("========= %+v", res)

	// in := &GoodsInfoRequest{
	// 	Page:        1,
	// 	RequestID:   "1",
	// 	ChannelType: 1,
	// }

	// res, err := c.QueryGoodsList(in)
	// if err != nil {
	// 	t.Errorf("%+v", err)
	// 	return
	// }
	// t.Logf("========= %+v", res)
}
