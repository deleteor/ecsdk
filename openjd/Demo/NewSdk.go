package Demo

import (
	JdunionSdk "git.hp6h2.cn/gobin/SDK/openjd"
)

var NewJDSdk JdunionSdk.JDSDKAPI

func init() {
	JdunionSdk.New(APPKEY, APPSECRET)
	NewJDSdk = JdunionSdk.JDSDKConfig
}
