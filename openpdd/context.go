package pdd

// Context 请求主体
type Context struct {
	ClientID     string
	ClientSecret string
	RetryTimes   int
	Debug        bool
}

// NewContext .
func NewContext(cfg *Config) *Context {
	return &Context{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RetryTimes:   cfg.RetryTimes,
		Debug:        cfg.Debug,
	}
}
