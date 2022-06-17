package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

/*
因为go-playground/validator 默认的错误信息是英文，但我们的错误信息不一定是用的英文，有可能要简体中文
通过中间件配合语言包的方式去实现简单的国际化需求
*/

func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")
		//global.Logger.Debugger("locale 是个啥: %v", locale)
		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			c.Set("trans", trans)
		}

		c.Next()
	}
}
