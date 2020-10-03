package libs

import (
	"errors"

	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/labstack/echo"

	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func DecodeValidate(v interface{}, c echo.Context) error {
	if err := c.Bind(v); err != nil {
		return err
	}
	if err := ValidateWithTraslator(v); err != nil {
		return err
	}
	return nil
}
func Validate(s interface{}) error {
	validate = validator.New()
	err := validate.Struct(s)
	if err != nil {
		return err
	}
	return nil
}

func ValidateWithTraslator(s interface{}) error {
	es := es.New()
	trans, _ := ut.New(es, es).GetTranslator("es")
	validate = validator.New()
	es_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(s)
	errStr := ""
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errStr += "\n" + e.Translate(trans)
		}
		return errors.New(errStr)
	}
	return nil
}
