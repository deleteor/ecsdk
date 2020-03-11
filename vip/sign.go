package vip

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// HMACMD5 hmac-md5加密
func (c *Client) HMACMD5(params map[string]interface{}, parmStr string) string {
	str := c.SortParamsToStr(params) + parmStr
	fmt.Printf("加密串 %s \n", str)
	fmt.Printf("key %s \n", c.AppSecret)
	mac := hmac.New(md5.New, []byte(c.AppSecret))
	mac.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)))
	// return hex.EncodeToString(mac.Sum(nil))
}
