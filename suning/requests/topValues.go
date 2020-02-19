package requests

import (
	"encoding/json"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

// TopValues 自定义的请求参数类型
type TopValues map[string]interface{}

// Add 添加新的请求参数
func (vs TopValues) Add(k string, v interface{}) {
	vs[k] = getString(v)
}

// GetSortedKeys 获取已排序的KEY值
func (vs TopValues) GetSortedKeys() []string {
	return sortKeys(vs)
}

// Encode 编码请求参数
func (vs TopValues) Encode() string {
	if vs == nil {
		return ""
	}
	var buf strings.Builder
	keys := make([]string, 0, len(vs))
	for k := range vs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		val := vs[k]
		keyEscaped := url.QueryEscape(k)
		buf.WriteString(keyEscaped + "=")
		buf.WriteString(url.QueryEscape(getString(val)))
	}
	return buf.String()
}

func getString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int:
		return strconv.Itoa(v)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := json.Marshal(v)
		return string(bytes)
	}
	return ""
}

// ValidateRequired 校验必填项
func (vs TopValues) ValidateRequired(k string) {
	if v, ok := vs[k]; !ok || len(getString(v)) == 0 {
		panic("field: " + k + " is required.")
	}
}

// sortKeys 内部排序
func sortKeys(vs TopValues) []string {
	sortedKeys := make([]string, 0, len(vs))
	for k := range vs {
		sortedKeys = append(sortedKeys, k)
	}
	sort.Strings(sortedKeys)
	return sortedKeys
}
