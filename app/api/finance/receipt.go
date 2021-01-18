package finance

import (
	"vss/app/service"

	"github.com/gin-gonic/gin"
)

func GetMergeList(c *gin.Context) {
	//验证数据
	var reqData service.GetReceiptAndReturnInput
	if err := c.ShouldBind(&reqData); err != nil {
		service.ResponseData{Message: err.Error()}.Json(c)
		return
	}
	////获取收货退货数据
	list, total := service.GetReceiptAndReturnData(reqData)
	service.ResponseData{Data: service.ResponseListData{List: list, Total: total}, Message: "2323"}.Json(c)
	return
}
