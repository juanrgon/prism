package prism

import (
	"github.com/juanrgon/prism/ansi"
	"github.com/juanrgon/prism/utils"
)

// InBlack returns the input string with black text
func InBlack(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.BlackText))
}

// InRed returns the input string with red text
func InRed(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.RedText))
}

// InGreen returns the input string with green text
func InGreen(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.GreenText))
}

// InYellow returns the input string with yellow text
func InYellow(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.YellowText))
}

// InBlue returns the input string with blue text
func InBlue(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.BlueText))
}

// InMagenta returns the input string with magenta text
func InMagenta(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.MagentaText))
}

// InCyan returns the input string with cyan text
func InCyan(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.CyanText))
}

// InWhite returns the input string with white text
func InWhite(s string) DecoratedString {
	return DecoratedString(utils.ChangeTextColor(s, ansi.WhiteText))
}
