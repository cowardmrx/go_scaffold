package global

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"go_scaffold/pkg/validator_rules"
	"reflect"
	"strings"
)

var ValidatorTrans ut.Translator
var Validate *validator.Validate

//	@method InitValidatorTranslator
//	@description: 加载validator验证器的翻译机
func InitValidatorTranslator() {
	//修改gin框架中的Validator属性，实现自定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册一个获取json tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("label"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT)

		// locale 通常取决于 http 请求头的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
		ValidatorTrans, ok = uni.GetTranslator("zh")
		if !ok {
			panic("uni.GetTranslator failed")
		}

		err := OverrideTranslator(v, ValidatorTrans)
		if err != nil {
			panic("override already validator rule message failed:" + err.Error())
		}

		loadSelfValidatorRules(v)

		loadSelfValidatorMessage(v, ValidatorTrans)

		Validate = v
	}

}

//	OverrideTranslator
//	@description: 覆盖原有规则的翻译
//	@param v *validator.Validate
//	@param translator ut.Translator
//	@return error
func OverrideTranslator(v *validator.Validate, translator ut.Translator) error {
	// 添加额外翻译 required_with
	_ = v.RegisterTranslation("required_with", translator, func(ut ut.Translator) error {
		return ut.Add("required_with", "{0}为必填字段!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_with", fe.Field())
		return t
	})

	// required_with_all
	_ = v.RegisterTranslation("required_with_all", translator, func(ut ut.Translator) error {
		return ut.Add("required_with_all", "{0} 为必填字段!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_with_all", fe.Field())
		return t
	})

	// required_without
	_ = v.RegisterTranslation("required_without", translator, func(ut ut.Translator) error {
		return ut.Add("required_without", "{0} 为必填字段!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_without", fe.Field())
		return t
	})

	// required_without_all
	_ = v.RegisterTranslation("required_without_all", translator, func(ut ut.Translator) error {
		return ut.Add("required_without_all", "{0} 为必填字段!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_without_all", fe.Field())
		return t
	})

	// unique
	_ = v.RegisterTranslation("unique", translator, func(ut ut.Translator) error {
		return ut.Add("unique", "{0}存在重复值", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("unique", fe.Field())
		return t
	})

	err := zhTranslations.RegisterDefaultTranslations(v, translator)
	if err != nil {
		return err
	}

	return nil
}

//	loadSelfValidatorRules
//	@description: 加载自定义校验规则
//	@param validate *validator.Validate
func loadSelfValidatorRules(validate *validator.Validate) {
	// 注册国内手机号校验规则
	err := validate.RegisterValidation("ZHPhone", validator_rules.ZHPhone)

	if err != nil {
		panic("register diy validator ZHPhone rules failed:" + err.Error())
	}

	// 注册校验参数类型是否为字符串规则
	err = validate.RegisterValidation("string", validator_rules.IsString)

	if err != nil {
		panic("register diy validator isString rules failed:" + err.Error())
	}

}

//	loadSelfValidatorMessage
//	@description: 加载自定义校验规则的翻译
//	@param validate *validator.Validate
//	@param trans ut.Translator
func loadSelfValidatorMessage(validate *validator.Validate, trans ut.Translator) {
	// 注册国内手机号校验规则的中文错误信息提示
	err := validate.RegisterTranslation("ZHPhone", trans, func(ut ut.Translator) error {
		return ut.Add("ZHPhone", "手机号格式错误", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("ZHPhone", fe.Field(), fe.Field())
		return t
	})
	if err != nil {
		panic("register diy validator translator failed:" + err.Error())
	}

	// 注册校验参数类型是否为字符串规则的中文错误信息提示
	err = validate.RegisterTranslation("string", trans, func(ut ut.Translator) error {
		return ut.Add("string", "参数类型非字符串格式", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("string", fe.Field(), fe.Field())
		return t
	})
	if err != nil {
		panic("register diy validator translator failed:" + err.Error())
	}
}
