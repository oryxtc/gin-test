package service

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

/**
struct转map
*/
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// recover错误，转string
func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	case int:
		return strconv.Itoa(v)
	case nil:
		return ""
	default:
		return v.(string)
	}
}

func GetLastErr(c *gin.Context) error {
	return c.Errors.Last().Err
}
