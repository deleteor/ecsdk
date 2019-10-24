package opentb

// FilesRequest 请求参数
type FilesRequest struct {
	// StartDate Date	true	2017-11-11 00:00:00	搜索开始时间
	StartDate string `json:"start_date" form:"start_date"`
	// EndDate Date	true	2017-11-12 00:00:00	搜索结束时间
	EndDate string `json:"end_date" form:"end_date"`
	// Status Number	false	1	下载链接状态。1:未下载。2:已下载
	Status int `json:"status" form:"status"`
}

// GetFiles 业务文件获取
func (c *Client) GetFiles(in *FilesRequest, cfg *Config) (map[string]interface{}, error) {
	c.SetMethod(FilesGet)

	params := GetParams(in)
	result, err := c.DoRequest(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
