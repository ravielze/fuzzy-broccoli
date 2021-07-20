package v10

import (
	"errors"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	v10 "github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/ravielze/oculi/validator"
	"github.com/ravielze/oculi/validator/v10/custom"
)

type (
	impl struct {
		instance *v10.Validate
		trans    ut.Translator
	}

	RegisterValidator struct {
		tag string
		vr  validator.ValidatorRegisterable
	}
)

var (
	defaultValidator = []RegisterValidator{
		{"after_now", custom.AfterNow},
		{"before_now", custom.BeforeNow},
		{"base36", custom.Base36},
	}
)

func New() (validator.Validator, error) {
	langEn := en.New()
	langId := id.New()
	uni := ut.New(langEn, langEn, langId)
	trans, _ := uni.GetTranslator("en")

	validate := v10.New()
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}

	instance := &impl{instance: validate, trans: trans}
	if err := instance.installDefaultValidator(); err != nil {
		return nil, err
	}
	return instance, nil
}

func (i *impl) AddTranslation(tag string, format string, extraParams ...string) error {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, format, false)
	}
	transFn := func(ut ut.Translator, fe v10.FieldError) string {
		x := []string{fe.Field(), fe.Param()}
		x = append(x, extraParams...)
		t, err := ut.T(fe.Tag(), x...)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}
	return i.instance.RegisterTranslation(tag, i.trans, registerFn, transFn)
}

func (i *impl) Translator() *ut.Translator {
	return &i.trans
}

func (i *impl) installDefaultValidator() error {
	for _, v := range defaultValidator {
		if err := i.Register(v.tag, v.vr); err != nil {
			return err
		}
	}
	return nil
}

func (i *impl) Validate(object interface{}) error {
	return i.instance.Struct(object)
}

func (i *impl) ValidateVar(obj interface{}, tag string) error {
	return i.instance.Var(obj, tag)
}

func (i *impl) RegisterValidation(tag string, fn interface{}) {
	if fnConv, ok := fn.(v10.Func); !ok {
		panic("fn is not v10.Func")
	} else {
		i.instance.RegisterValidation(tag, fnConv)
	}
}

func (i *impl) RegisterCustomTypeFunc(fn validator.CustomTypeFunc, types ...interface{}) {
	i.instance.RegisterCustomTypeFunc(v10.CustomTypeFunc(fn), types...)
}

func (i *impl) RegisterStructValidation(fn interface{}, types ...interface{}) {
	if fnConv, ok := fn.(v10.StructLevelFunc); !ok {
		panic("fn is not v10.StructLevelFunc")
	} else {
		i.instance.RegisterStructValidation(fnConv, types...)
	}
}

func (i *impl) Register(tag string, vr validator.ValidatorRegisterable) error {
	fn, format, extraParams := vr()
	fnConv, ok := fn.(v10.Func)
	if !ok {
		return errors.New("fn is not v10.Func")
	}

	if err := i.instance.RegisterValidation(tag, fnConv); err != nil {
		return err
	}
	return i.AddTranslation(tag, string(format), extraParams...)
}