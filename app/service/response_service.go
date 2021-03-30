package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var httpCode = http.StatusOK //默认返回200状态

type ResponseData struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseListData struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
}

/**
设置httpCode
*/
func (r ResponseData) SetHttpCode(code int) ResponseData {
	httpCode = code
	return r
}

/**
返回Json
*/
func (r ResponseData) Json(c *gin.Context) {
	res := ResponseData{Code: r.Code, Data: r.Data, Message: r.Message}
	c.JSON(httpCode, res)
}

/*
实现error 接口
*/
func (r ResponseData) Error() string {
	return r.Message
}
