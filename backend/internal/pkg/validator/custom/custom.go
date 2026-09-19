// Package custom contains type ValidTag that represents a custom validation tag
// and ValidTag instances.
package custom

import (
	"fmt"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// ValidTag represents a custom validation tag.
// Translations map must contain "en" translation.
type ValidTag struct {
	Tag string

	validFunc    func(fl validator.FieldLevel) bool
	translations translations
}

// translations represents all tag translations for different locales (map[locale]translation).
type translations map[string]string

// get returns translated msg by the given locale.
// If locale was not found it returns default (en) translation.
func (t translations) get(locale string) string {
	trans, ok := t[locale]
	if !ok {
		return t["en"]
	}
	return trans
}

// Register registers custom tag validation.
func (v *ValidTag) Register(inst *validator.Validate) error {
	if err := inst.RegisterValidation(v.Tag, v.validFunc); err != nil {
		return fmt.Errorf("register validation: %w", err)
	}
	return nil
}

// RegisterWithTrans registers custom tag validation
// with error translation uses the given locale.
func (v *ValidTag) RegisterWithTrans(inst *validator.Validate,
	transInst ut.Translator, transLocale string) error {

	// validation
	if err := inst.RegisterValidation(v.Tag, v.validFunc); err != nil {
		return fmt.Errorf("register validation: %w", err)
	}

	// translation
	err := inst.RegisterTranslation(v.Tag, transInst,
		func(ut ut.Translator) error {
			return ut.Add(v.Tag, v.translations.get(transLocale), true)
		},
		func(ut ut.Translator, _ validator.FieldError) string {
			t, _ := ut.T(v.Tag) // nolint:errcheck // tag v.tag was registered
			return t
		},
	)
	if err != nil {
		return fmt.Errorf("register translation: %w", err)
	}
	return nil
}
