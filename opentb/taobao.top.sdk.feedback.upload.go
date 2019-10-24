package opentb

// TopSdkFeedbackUploadRequest sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
type TopSdkFeedbackUploadRequest struct {
	// Type true	1	1、回传加密信息
	Type string `json:"type" form:"type"`
	// Content false	{"key":"value"}	具体内容，json形式
	Content string `json:"content" form:"content"`
}

// UploadTopSdkFeedback sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
func (c *Client) UploadTopSdkFeedback(in *TopSdkFeedbackUploadRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(TopSdkFeedbackUpload)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
