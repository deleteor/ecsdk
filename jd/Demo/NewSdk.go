package Demo

import (
	JdunionSdk "ecsdk/jd"
)

var NewJDSdk JdunionSdk.JDSDKAPI

func init() {
	JdunionSdk.New(APPKEY, APPSECRET)
	NewJDSdk = JdunionSdk.JDSDKConfig
}
