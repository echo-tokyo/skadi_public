package custom

import (
	"skadi/backend/internal/pkg/password"

	"github.com/go-playground/validator/v10"
)

var StrongPasswd = &ValidTag{
	Tag: "strong-passwd",
	validFunc: func(fl validator.FieldLevel) bool {
		return password.Strong(fl.Field().String())
	},
	translations: map[string]string{
		"en": "password too weak",
		"ru": "слабый пароль",
	},
}
