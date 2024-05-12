package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translate "github.com/go-playground/validator/v10/translations/en"
	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translate.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *resterr.RestErr {

	// Validate invalid fild type
	var jsonErr *json.UnmarshalTypeError

	// Error of the validation
	var jsonValidationError validator.ValidationErrors

	// if errors is the same type validation_err with jsonErr
	if errors.As(validation_err, &jsonErr) {
		return resterr.NewBadRequestError("Invalid field type")

		// if errors is the same type validation_err with jssonValidationError
	} else if errors.As(validation_err, &jsonValidationError) {
		// array causes
		errorsCauses := []resterr.Causes{}

		// _ -> ignore the first param
		for _, e := range validation_err.(validator.ValidationErrors) {
			causse := resterr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, causse)
		}

		return resterr.NewBadRequestValidationError("Some fields are invalid", errorsCauses)
	} else {
		return resterr.NewBadRequestError("Error trying to convert fields")
	}
}
