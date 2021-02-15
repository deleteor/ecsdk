package tb

// Config 淘宝配置
type Config struct {
	// 好单库密钥
	HaodankuApikey string `yaml:"haodankuApikey"`
	// 淘宝第一账号
	AppKey    string `yaml:"appKey"`
	SecretKey string `yaml:"secretKey"`
	// 淘宝渠道账号
	ChannelAppKey    string `yaml:"channelAppKey"`
	ChannelSecretKey string `yaml:"channelSecretKey"`
	// 老的mm_170480121_78100180  即将更换 mm_314740118_361750003
	UnionID int `yaml:"unionID"`
	// 导购位
	SiteID int `yaml:"siteID"`
	// 默认pid
	DeafaultPid string `yaml:"deafaultPid"`
	// 渠道Pid
	RelationPid      string `yaml:"relationPid"`
	NoChannelUnionID []int  `yaml:"noChannelUnionID"`
	// 淘宝渠道邀请码
	InviteCode string `yaml:"inviteCode"`
	BlackList  int    `yaml:"blackList"`
}
