package base

import (
	"fmt"
	"reflect"
	"review/internal/models"

	"github.com/gin-gonic/gin"
)

func User(ctx *gin.Context) *models.User {
	table := &models.User{}
	if user, ok := ctx.Get("user"); ok {
		// 使用反射来获取和设置字段
		userValue := reflect.ValueOf(user)
		if userValue.Kind() == reflect.Ptr {
			userValue = userValue.Elem()
		}

		tableValue := reflect.ValueOf(table).Elem()

		for i := 0; i < userValue.NumField(); i++ {
			fieldName := userValue.Type().Field(i).Name
			fieldValue := userValue.Field(i)

			tableField := tableValue.FieldByName(fieldName)
			if tableField.IsValid() && tableField.CanSet() {
				tableField.Set(fieldValue)
			}
		}

	} else {
		fmt.Println("上下文中没有找到用户信息")
		return nil
	}
	return table
}
