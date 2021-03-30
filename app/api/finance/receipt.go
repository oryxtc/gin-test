package finance

import (
	"vss/app/service"

	"github.com/gin-gonic/gin"
)

func GetMergeList(c *gin.Context) {
	//验证数据
	var reqData service.GetReceiptAndReturnInput
	if err := c.ShouldBind(&reqData); err != nil {
		panic(err)
	}
	////获取收货退货数据
	list, total, err := service.GetReceiptAndReturnData(reqData)
	if err != nil {
		panic(err)
	}
	service.ResponseData{Data: service.ResponseListData{List: list, Total: total}}.Json(c)
	return
}
