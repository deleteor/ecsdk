package pdd

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func createSign(signStr string) string {
	h := md5.New()
	h.Write([]byte(signStr))
	cipherStr := h.Sum(nil)
	return strings.ToUpper(hex.EncodeToString(cipherStr))
}
