package Demo

import (
	"encoding/json"
	"fmt"
	"testing"

	JdunionSdk "ecsdk/jd"
)

func TestCutomInterfaceGoods(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys2 := NewJDSdkCutom.NewFunc(ParamJson)
	t.Log(categorys2)
}

//查询分类
func TestCateGorys(t *testing.T) {
	ParamJson := "{\"req\":{\"parentId\":0,\"grade\":0}}"
	categorys := NewJDSdk.GetCategoryList(ParamJson)
	t.Log(categorys)
}

//查询订单
func TestOrder(t *testing.T) {
	ParamStruct := JdunionSdk.OrderParam{}
	ParamStruct.OrderReq.Time = "2019072010"
	ParamStruct.OrderReq.PageNo = 1
	ParamStruct.OrderReq.PageSize = 10
	ParamStruct.OrderReq.Type = 1
	bytes, _ := json.Marshal(ParamStruct)
	fmt.Println(string(bytes))

	orders := NewJDSdk.GetOrders(string(bytes))
	fmt.Println(orders)

}

//京粉
func TestJF(t *testing.T) {
	ParamJF := JdunionSdk.ParamJFReq{}
	ParamJF.GoodsReq.PageSize = 20
	ParamJF.GoodsReq.EliteId = 2
	ParamJF.GoodsReq.PageIndex = 1
	bytes, _ := json.Marshal(ParamJF)
	fen := NewJDSdk.GetGoodsJFen(string(bytes))
	t.Log(fen)
}
