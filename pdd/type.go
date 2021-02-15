package pdd

// MethodType ddk接口
type MethodType string

// 多多客
const (
	DDKThemeListGet MethodType = "pdd.ddk.theme.list.get"
	// 生成红包推广链接
	DDKRPPromURLGenerate         MethodType = "pdd.ddk.rp.prom.url.generate"
	DDKOrderListIncrementGet     MethodType = "pdd.ddk.order.list.increment.get"
	DDKColorOrderIncrementGet    MethodType = "pdd.ddk.color.order.increment.get"
	DDKGoodsDetail               MethodType = "pdd.ddk.goods.detail"
	DDKGoodsSearch               MethodType = "pdd.ddk.goods.search"
	DDKGoodsPidQuery             MethodType = "pdd.ddk.goods.pid.query"
	DDKGoodsPidGenerate          MethodType = "pdd.ddk.goods.pid.generate"
	DDKGoodsPromotionURLGenerate MethodType = "pdd.ddk.goods.promotion.url.generate"
	DDKLotteryURLGen             MethodType = "pdd.ddk.lottery.url.gen"
	DDKCMSPromURLGenerate        MethodType = "pdd.ddk.cms.prom.url.generate"
	//多多客获取爆款排行商品接口
	DDKTopGoodsListQuery MethodType = "pdd.ddk.top.goods.list.query"
	//多多进宝转链接口
	DDKGoodsZsUnitURLGen MethodType = "pdd.ddk.goods.zs.unit.url.gen"

	//商品 API
	//商品标准类目接口
	GoodsCatsGet MethodType = "pdd.goods.cats.get"
	//查询商品标签列表
	GoodsOptGet MethodType = "pdd.goods.opt.get"
)
