package validator

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	entran "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validate(data any) (bool, map[string][]string) {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate = validator.New()
	validate.RegisterValidation("notblank", validators.NotBlank)
	entran.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(data)

	errors := make(map[string][]string)
	reflected := reflect.ValueOf(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			field, _ := reflected.Type().FieldByName(err.StructField())
			var name string

			if name = field.Tag.Get("json"); name == "" {
				name = strings.ToLower(err.StructField())
			}

			errors[name] = append(errors[name], err.Translate(trans))
		}
	}

	if len(errors) == 0 {
		return true, nil
	}

	return false, errors
}
