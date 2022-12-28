package text_validator

import "strings"

type textIsAbbaValidator struct{}

func NewIsAbbaTextValidator() *textIsAbbaValidator {
	return &textIsAbbaValidator{}
}

func (tv *textIsAbbaValidator) IsValid(tekst string) bool {
	return strings.HasPrefix(tekst, "abba")
}
