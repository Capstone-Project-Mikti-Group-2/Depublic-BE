package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

type FormValidator struct {
	validator *validator.Validate
}

func (fv *FormValidator) Validate(i interface{}) error {
	return fv.validator.Struct(i)
}

func NewFormValidator() *FormValidator {
	validate := validator.New(validator.WithRequiredStructEnabled())

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return &FormValidator{validate}
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				fields[err.Field()] = fmt.Sprintf("field %s harus di isi", err.Field())
			case "password":
				fields[err.Field()] = "Password harus mengandung setidaknya satu huruf besar dan nomor"
			case "oneof":
				fields[err.Field()] = fmt.Sprintf("field %s harus di isi dengan salah satu dari %s", err.Field(), err.Param())
			default:
				fields[err.Field()] = fmt.Sprintf("kesalahan pada %s dengan tag %s seharusnya %s ", err.Field(), err.Tag(), err.Param())
			}
		}
	}
	return fields
}
