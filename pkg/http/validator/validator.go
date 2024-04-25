package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

var Validator *validator.Validate
var Trans ut.Translator

func Init() {
	Validator = validator.New()
	//注册自定义验证器
	RegisterValidation(Validator)
	//注册一个获取json的自定义方法
	Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("form"), ",", 2)[0]
		if name == "-" {
			name = field.Name
		}
		return name
	})
}

func SetTranslator(locales ...string) {
	zhT := zh.New() // 中文翻译器
	enT := en.New() // 英文翻译器

	// 第一个参数是备用(fallback)的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(zhT, zhT, enT)

	// locale 通常取决于 http 请求头的 'Accept-Language'
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	Trans, _ = uni.FindTranslator(locales...)
	// 注册翻译器
	_ = zhTranslations.RegisterDefaultTranslations(Validator, Trans)

	//注册自定义翻译
	RegisterTranslation(Validator)
}
