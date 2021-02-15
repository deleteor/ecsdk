# suning
苏宁开放平台SDK  
 

```golang
package main

import (
	"fmt"
	"suning"
	"suning/requests"
	"suning/responses" 
)

 
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	appKey, appSecret := "156c3105338f79d12bebfdbb146f5280", "958b0e2bd727b762f9d0787588fa3acf"
	client, _ := suning.NewTopClient(appKey, appSecret)
    /*
	req := requests.NewNetallianceCommoditydetailQueryGetRequest()
	req.SetCityCode("025")
	req.SetPicWidth("640")
	req.SetPicHeight("640")
	req.SetCommodityStr("107053373560-0000000000")
	res := responses.NetallianceCommoditydetailQueryGetResponse{}
	err := client.Do(req, &res)
	checkErr(err)
	fmt.Println(len(res.ResponseContent.ResponseBody.QueryCommoditydetail))
	*/

    /*
	req2 := requests.NewNetallianceOrderQueryGetRequest()
	req2.SetPageNo(1)
	req2.SetPageSize(2)
    req2.SetStartTime("2020-02-14 11:30:30")
    req2.SetEndTime("2020-02-14 11:50:30")
	req2.SetOrderLineStatus(0)
	res2 := responses.NetallianceOderQueryGetResponse{}
	err2 := client.Do(req2, &res2)
	checkErr(err2)
	result := res2.ResponseContent.ResponseBody.QueryOrder
	for _,v := range result {
		fmt.Println(v.OrderCode)
		fmt.Println(v.OrderDetail)
	}*/

	/*
    req := requests.NewNetallianceInverstmentcategoryidQueryGetRequest()
	res := responses.NetallianceInverstmentcategoryidQueryGetResponse{}
	err := client.Do(req, &res)
	checkErr(err)
	cateList := res.ResponseContent.ResponseBody.ResponseCategory.CategoryList
	for _,v := range cateList  {
		fmt.Println(v.Name,v.ID)
	}*/

	/*
	req := requests.NewNetallianceInverstmentcommodityQueryGetRequest()
	req.SetPageSize(10)
	req.SetPageIndex(2)
	req.SetCategoryId("C1")

	res := responses.NetallianceInverstmentcommodityQueryGetResponse{}
	err := client.Do(req, &res)
	checkErr(err)
	cateList := res.ResponseContent.ResponseBody
	*/

   /*
	req := requests.NewNetallianceCouponinfoQueryGetRequest()
	req.SetQuanUrl("https://quan.suning.com/lqzx_recommend.do?activityId=202002080012370155&activitySecretKey=LPXpXRZFiL6ilu7e8fjwFYF8")

	res := responses.NetallianceCouponinfoQueryGetResponse{}
	err := client.Do(req, &res)
	checkErr(err)
	info := res.ResponseContent.ResponseBody
	fmt.Println(info)
    */
   /*
	req := requests.NewNetallianceMorerecommendGetRequest()
	req.SetCommodityCode("181754507")
	req.SetSupplierCode("00")

	res := responses.NetallianceMorerecommendGetResponse{}
	err := client.Do(req, &res)
	checkErr(err)
	info := res.ResponseContent.ResponseBody
    */
   /*
	req := requests.NewNetallianceRecommendcommodityQueryGetRequest()
	req.SetPageSize(10)
	req.SetPageIndex(1)

	res := responses.NetallianceRecommendcommodityQueryGetResponse{}
	*/
    /*
	req := requests.NewNetallianceSearchcommodityQueryGetReuqest()
	req.SetPageSize(10)
	req.SetPageIndex(1)
	//req.SetKeyword("水果")
    req.SetSaleCategoryCode(500353)
	res := responses.NetallianceSearchcommodityQueryGetResponse{}
	*/

	req := requests.NewNetallianceAppletextensionlinkGetRequest()
	req.SetProductUrl("https://product.suning.com/0000000000/10553920376.html")
	req.SetQuanUrl("https://quan.suning.com/lqzx_recommend.do?activityId=202002080012370155&activitySecretKey=LPXpXRZFiL6ilu7e8fjwFYF8")
	req.SetSubUser("17624")
	res := responses.NetallianceAppletextensionlinkGetResponse{}
	err := client.Do(req, &res)
	checkErr(err)
	info := res.ResponseContent.ResponseBody.GetAppletextensionlink
	fmt.Println(info.AppId,info.AppletExtendUrl) 
}

```