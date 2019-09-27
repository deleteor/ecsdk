package openjd

import (
	"io/ioutil"
	"net/http"
)

// HttpGet 请求
func HttpGet(url string) ([]byte, error) {
	req, _ := http.NewRequest(CUSTOMREQMETHODGET, url, nil)
	//设置请求头部信息
	//req.Header.Add("Authorization", q.Token)
	//发送请求
	response, _ := http.DefaultClient.Do(req)
	//关闭请求
	defer Close(response)
	//解析返回结果
	bytes, err := ioutil.ReadAll(response.Body)
	return bytes, err
}

// Close 关闭
func Close(response *http.Response) {

	e := response.Body.Close()
	if e != nil {
		panic(e)
	}
}
