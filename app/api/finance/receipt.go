package finance

import (
	"net/http"
	"vss/app/service"

	"github.com/gin-gonic/gin"
)

func GetMergeList(c *gin.Context) {
	//验证数据
	var reqData service.GetReceiptAndReturnInput
	if err := c.Bind(&reqData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//获取收货退货数据
	data, total := service.GetReceiptAndReturnData(reqData)
	c.JSON(200, gin.H{"data": data, "total": total})
}
