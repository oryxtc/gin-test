package middleware

import (
	"log"
	"net/http"
	"vss/app/service"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				log.Printf("panic: %v\n", r)
				c.JSON(http.StatusOK, gin.H{
					"code": 400,
					"msg":  service.ErrorToString(r),
					"data": []interface{}{},
				})
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		c.Next()
	}
}
