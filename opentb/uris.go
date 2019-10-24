package opentb

// TaoBaoURI 淘宝URI
type TaoBaoURI string

const (
	// AppIPGet 获取ISV发起请求服务器IP
	AppIPGet TaoBaoURI = "taobao.appip.get"
	// FilesGet 业务文件获取
	FilesGet TaoBaoURI = "taobao.files.get"
	// HTTPDnsGet Top Dns 配置
	HTTPDnsGet TaoBaoURI = "taobao.httpdns.get"
	// KFCKeywordSearch 关键词过滤匹配
	KFCKeywordSearch TaoBaoURI = "taobao.kfc.keyword.search"
	// OpenuidGet 获取授权账号对应的OpenUid
	OpenuidGet TaoBaoURI = "taobao.openuid.get"
	// OpenuidChange 淘宝openUid 转换
	OpenuidChange TaoBaoURI = "taobao.openuid.change"
	// OpenuidGetBymixnick 通过mixnick转换openuid
	OpenuidGetBymixnick TaoBaoURI = "taobao.openuid.get.bymixnick"
	// OpenuidGetBytrade 通过订单获取对应买家的openUID
	OpenuidGetBytrade TaoBaoURI = "taobao.openuid.get.bytrade"
	// TbkCouponGet 阿里妈妈推广券信息查询
	TbkCouponGet TaoBaoURI = "taobao.tbk.coupon.get"
	// TbkItemInfoGet 淘宝客-公用-淘宝客商品详情查询(简版)
	TbkItemInfoGet TaoBaoURI = "taobao.tbk.item.info.get"
	// TbkPrivilegeGet  淘宝客-服务商-单品券高效转链
	TbkPrivilegeGet TaoBaoURI = "taobao.tbk.privilege.get"
	// TbkScActivitylinkToolget 淘宝客-服务商-官方活动转链
	TbkScActivitylinkToolget TaoBaoURI = "taobao.tbk.sc.activitylink.toolget"
	// TbkScAdzoneCreate 淘宝客-服务商-创建推广者位
	TbkScAdzoneCreate TaoBaoURI = "taobao.tbk.sc.adzone.create"
	// TbkScGroupchatCreate 淘宝客-服务商-手淘群创建
	TbkScGroupchatCreate TaoBaoURI = "taobao.tbk.sc.groupchat.create"
	// TbkScGroupchatGet 淘宝客-服务商-手淘群查询
	TbkScGroupchatGet TaoBaoURI = "taobao.tbk.sc.groupchat.get"
	// TbkScGroupchatMessageSend 淘宝客-服务商-手淘群发单
	TbkScGroupchatMessageSend TaoBaoURI = "taobao.tbk.sc.groupchat.message.send"
	// TbkScInvitecodeGet 淘宝客-公用-私域用户邀请码生成
	TbkScInvitecodeGet TaoBaoURI = "taobao.tbk.sc.invitecode.get"
	// TbkScMaterialOptional 淘宝客-服务商-物料搜索
	TbkScMaterialOptional TaoBaoURI = "taobao.tbk.sc.material.optional"
	// TbkScOptimusMaterial 淘宝客-服务商-物料精选
	TbkScOptimusMaterial TaoBaoURI = "taobao.tbk.sc.optimus.material"
	// TbkScOrderDetailsGet 淘宝客-服务商-所有订单查询
	TbkScOrderDetailsGet TaoBaoURI = "taobao.tbk.sc.order.details.get"
	// TbkScOrderGet .
	TbkScOrderGet TaoBaoURI = "taobao.tbk.sc.order.get"
	// TbkScPublisherInfoGet 淘宝客-公用-私域用户备案信息查询
	TbkScPublisherInfoGet TaoBaoURI = "taobao.tbk.sc.publisher.info.get"
	// TbkScPublisherInfoSave 淘宝客-公用-私域用户备案
	TbkScPublisherInfoSave TaoBaoURI = "taobao.tbk.sc.publisher.info.save"
	// TbkSpreadGet 淘宝客-公用-长链转短链
	TbkSpreadGet TaoBaoURI = "taobao.tbk.spread.get"
	// TbkTpwdCreate 淘宝客-公用-淘口令生成
	TbkTpwdCreate TaoBaoURI = "taobao.tbk.tpwd.create"
	// TimeGet 获取淘宝系统当前时间
	TimeGet TaoBaoURI = "taobao.time.get"
	// TopAuthTokenCreate 获取Access Token 用户通过code换获取access_token，https only
	TopAuthTokenCreate TaoBaoURI = "taobao.top.auth.token.create"
	// TopAuthTokenRefresh 刷新Access Token 根据refresh_token重新生成token
	TopAuthTokenRefresh TaoBaoURI = "taobao.top.auth.token.refresh"
	// TopIpoutGet 获取开放平台出口IP段
	TopIpoutGet TaoBaoURI = "taobao.top.ipout.get"
	// TopSdkFeedbackUpload sdk回调客户端基本信息到开放平台，用于做监控之类，有助于帮助isv监控系统稳定性
	TopSdkFeedbackUpload TaoBaoURI = "taobao.top.sdk.feedback.upload"
	// TopSecretGet 获取TOP通道解密秘钥
	TopSecretGet TaoBaoURI = "taobao.top.secret.get"
	// TopSecretRegister 提供给isv注册非淘系账号秘钥，isv依赖sdk自主加、解密
	TopSecretRegister TaoBaoURI = "taobao.top.secret.register"
	// WirelessShareTpwdQuery 查询解析淘口令
	WirelessShareTpwdQuery TaoBaoURI = "taobao.wireless.share.tpwd.query"
)
