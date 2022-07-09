package validat

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"time"
)

func dateFormat(fl validator.FieldLevel) bool {
	if _, ok := fl.Field().Interface().(time.Time); ok {
		return false
	} else {
		return false
	}
}

// 	需要在main里调用validat.RegisterValidator()

func RegisterValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//绑定第一个参数是验证的函数第二个参数是自定义的验证函数
		v.RegisterValidation("dateFormat", dateFormat)
	}
}
