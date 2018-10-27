package prism

import (
	"github.com/juanrgon/prism/utils"
)

// Underlined returns the input string underlined
func Underlined(s string) DecoratedString {
	return DecoratedString(utils.Underline(s))
}

// InBold returns the input string in bold
func InBold(s string) DecoratedString {
	return DecoratedString(utils.Embolden(s))
}
