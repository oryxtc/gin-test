package middleware

import (
	"net/http"
	"vss/app/service"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//获取参数值
				var code int
				var data interface{}
				switch v := r.(type) {
				case service.ResponseData:
					if v.Code == 0 {
						v.Code = 400
					}
					code = v.Code
					data = v.Data
				default:
					code = 400
					data = nil
				}
				//打印错误堆栈信息
				c.JSON(http.StatusOK, gin.H{
					"code": code,
					"msg":  service.ErrorToString(r),
					"data": data,
				})
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		c.Next()
	}
}
