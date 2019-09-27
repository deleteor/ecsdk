package opentb

import (
	"crypto/hmac"
	"crypto/md5"
	"fmt"
)

// 签名算法
/*
* API算法的URL：http://open.taobao.com/docs/doc.htm?spm=a219a.7395905.0.0.hsp22E&articleId=101617&docType=1&treeId=1
*
 */

// SignMD5 MD5 加密
func (c *Client) SignMD5(params map[string]interface{}) string {
	str := fmt.Sprintf("%s%s%s", c.AppSecret, c.SortParamsToStr(params), c.AppSecret)
	return fmt.Sprintf("%X", md5.Sum([]byte(str)))
}

// SignHMAC HMAC 加密
func (c *Client) SignHMAC(params map[string]interface{}) string {
	str := c.SortParamsToStr(params)
	mac := hmac.New(md5.New, []byte(c.AppSecret))
	mac.Write([]byte(str))
	return fmt.Sprintf("%X", mac.Sum(nil))
}
