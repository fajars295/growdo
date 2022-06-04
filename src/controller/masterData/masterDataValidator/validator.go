package masterDatavalidator

import (
	"growdo/src/model"
	"strings"

	"github.com/go-playground/locales/id"
	ud "github.com/go-playground/universal-translator"
	en_translations "github.com/go-playground/validator/v10/translations/id"

	"github.com/go-playground/validator/v10"
)

func ValidatorCreate(r *model.MasterData) string {
	en := id.New()
	uni := ud.New(en, en)
	trans, _ := uni.GetTranslator("id")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)
	err := validate.Struct(r)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		var errors []string
		for _, e := range errs {
			errors = append(errors, e.Translate(trans))
		}
		return strings.Replace(strings.Join(errors, ","), "_", " ", -1)
	}

	return ""
}
