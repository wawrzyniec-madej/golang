package text_validator

import "strings"

type textHasExclamationValidator struct{}

func NewHasExclamationTextValidator() *textHasExclamationValidator {
	return &textHasExclamationValidator{}
}

func (tv *textHasExclamationValidator) IsValid(tekst string) bool {

	return strings.HasSuffix(tekst, "!")
}
