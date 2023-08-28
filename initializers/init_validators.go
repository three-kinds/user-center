package initializers

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/three-kinds/user-center/utils/generic_utils/validate_utils"
	"log"
	"reflect"
	"strings"
)

var UT *ut.UniversalTranslator

func registerCustomValidations(engine *validator.Validate) {
	err := engine.RegisterValidation("phone_number", func(fl validator.FieldLevel) bool {
		return validate_utils.IsPhoneNumber(fl.Field().String())
	})
	if err != nil {
		log.Panicln("register `phone_number` validation failed")
	}
}

func registerCustomZhTranslations(engine *validator.Validate, zhTranslator ut.Translator) {
	err := engine.RegisterTranslation(
		"phone_number", zhTranslator,
		registrationFunc("phone_number", "{0}必须是一个有效的手机号码", false),
		translateFunc,
	)
	if err != nil {
		log.Panicln("register `phone_number` translation failed")
	}
}

func InitValidators() {
	engine, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		log.Panicln(`validator Engine init error`)
	}

	registerCustomValidations(engine)

	// 注册一个获取json tag的自定义方法，返回错误字段使用 json tag 字段，而不是结构体字段名
	engine.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	enT := en.New()            // 英文翻译器
	zhT := zh.New()            // 中文翻译器
	UT = ut.New(enT, zhT, enT) // 第一个参数是备用语言环境，后面的参数是应该支持的语言环境

	// 注册中文翻译器
	if trans, ok := UT.GetTranslator("zh"); ok {
		if err := zhTranslations.RegisterDefaultTranslations(engine, trans); err != nil {
			log.Panicln(`RegisterTranslations("zh") error`)
		}
		registerCustomZhTranslations(engine, trans)
	} else {
		log.Panicln(`UT.GetTranslator("zh") error`)
	}

	// 注册英文翻译器
	if trans, ok := UT.GetTranslator("en"); ok {
		if err := enTranslations.RegisterDefaultTranslations(engine, trans); err != nil {
			log.Panicln(`RegisterTranslations("en") error`)
		}
	} else {
		log.Panicln(`UT.GetTranslator("en") error`)
	}

}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	t, err := ut.T(fe.Tag(), fe.Field())
	if err != nil {
		DefaultLogger.Warning(fmt.Sprintf("validator translate error: %#v", fe))
		return fe.(error).Error()
	}

	return t
}
