package controller

import (
	"fmt"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

// 1. No Global varibles, pass it explicitly where need
// 2. return nil when no errors

func InitTrans(locale string) (ut.Translator, error) {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return nil, fmt.Errorf("failed to get validator engine")
	}

	zhT := zh.New()
	enT := en.New()

	uni := ut.New(enT, zhT, enT)

	trans, ok := uni.GetTranslator(locale)
	if !ok {
		return nil, fmt.Errorf("failed to get translator for %s", locale)
	}

	var err error
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to register translator for %s", locale)
	}

	return trans, nil
}
