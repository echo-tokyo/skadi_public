package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ru"
	ut "github.com/go-playground/universal-translator"
	govalidator "github.com/go-playground/validator/v10"
	entranslation "github.com/go-playground/validator/v10/translations/en"
	rutranslation "github.com/go-playground/validator/v10/translations/ru"

	"skadi/backend/internal/pkg/validator/custom"
)

// Ensure tag validator implements interface.
var _ Validator = (*TagValidator)(nil)

// TagValidator implementats [Validator] interface.
// It validates structs by tags.
type TagValidator struct {
	validatorInst *govalidator.Validate
	transInst     ut.Translator
	transLocale   string
}

// NewTagValidator returns new instance of [TagValidator].
func NewTagValidator() (*TagValidator, error) {
	locale := "en"

	enTranslator := en.New()
	uni := ut.New(enTranslator, enTranslator)
	trans, _ := uni.GetTranslator(locale)

	validate := govalidator.New(govalidator.WithRequiredStructEnabled())
	err := entranslation.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, fmt.Errorf("register translation: %w", err)
	}

	validInst := &TagValidator{validate, trans, locale}
	if err := validInst.registerCustomValidFuncs(); err != nil {
		return nil, err
	}
	return validInst, nil
}

// NewRuTagValidator returns new instance of [TagValidator] with ru error translation.
func NewRuTagValidator() (*TagValidator, error) {
	locale := "ru"

	ruTranslator := ru.New()
	uni := ut.New(ruTranslator, ruTranslator)
	trans, _ := uni.GetTranslator(locale)

	validate := govalidator.New(govalidator.WithRequiredStructEnabled())
	err := rutranslation.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return nil, fmt.Errorf("register translation: %w", err)
	}

	validInst := &TagValidator{validate, trans, locale}
	if err := validInst.registerCustomValidFuncs(); err != nil {
		return nil, err
	}
	return validInst, nil
}

// registerCustomValidFuncs registers custom validate funcs for tag validator.
func (v TagValidator) registerCustomValidFuncs() error {
	// custom validation tags list
	customList := []*custom.ValidTag{
		custom.StrongPasswd,
	}

	// register all tags with translation
	var err error
	for _, validTag := range customList {
		err = validTag.RegisterWithTrans(v.validatorInst, v.transInst, v.transLocale)
		if err != nil {
			return fmt.Errorf("register tag %s: %w", validTag.Tag, err)
		}
	}
	return nil
}

// Validate validates given struct s (s is a pointer to struct).
func (v TagValidator) Validate(s any) error {
	err := v.validatorInst.Struct(s)
	if err == nil { // NOT err
		return nil
	}

	var validateErrors govalidator.ValidationErrors
	if !errors.As(err, &validateErrors) {
		return err
	}
	// handle error messages
	rawTranstaledMap := validateErrors.Translate(v.transInst)

	// sort out errors and concat them into string
	transtaledStringSlice := make([]string, 0, len(rawTranstaledMap))
	for _, v := range rawTranstaledMap {
		transtaledStringSlice = append(transtaledStringSlice, strings.ToLower(v))
	}

	return fmt.Errorf("%w: %s", ErrValidate, strings.Join(transtaledStringSlice, "\n"))
}
