package lib

import (
	"log"
	"unicode"
	"unicode/utf8"

	"github.com/go-playground/validator/v10"
)

// CutstomValidator :
type CustomValidator struct {
	Validator *validator.Validate
}

// Validate : Validate Data
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func NewCustomValidator() *CustomValidator {
	customValidator := &CustomValidator{Validator: validator.New()}
	customValidator.Init()
	return customValidator
}

func (cv *CustomValidator) Init() {
	err := cv.Validator.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		var (
			hasNumber      = false
			hasSpecialChar = false
			hasLetter      = false
			hasSuitableLen = false
		)

		password := fl.Field().String()

		if utf8.RuneCountInString(password) <= 30 && utf8.RuneCountInString(password) >= 8 {
			hasSuitableLen = true
		}

		for _, c := range password {
			switch {
			case unicode.IsNumber(c):
				hasNumber = true
			case unicode.IsPunct(c) || unicode.IsSymbol(c):
				hasSpecialChar = true
			case unicode.IsLetter(c) || c == ' ':
				hasLetter = true
			default:
				return false
			}
		}
		return hasNumber && hasSpecialChar && hasLetter && hasSuitableLen
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}
