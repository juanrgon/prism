package prism

import (
	"github.com/juanrgon/prism/ansi"
	"github.com/juanrgon/prism/utils"
)

// OnBlack returns the input string with a black background
func OnBlack(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.BlackBackground))
}

// OnRed returns the input string with a red background
func OnRed(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.RedBackground))
}

// OnGreen returns the input string with a green background
func OnGreen(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.GreenBackground))
}

// OnYellow returns the input string with a yellow background
func OnYellow(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.YellowBackground))
}

// OnBlue returns the input string with a blue background
func OnBlue(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.BlueBackground))
}

// OnMagenta returns the input string with a magenta background
func OnMagenta(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.MagentaBackground))
}

// OnCyan returns the input string with a cyan background
func OnCyan(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.CyanBackground))
}

// OnWhite returns the input string with a white background
func OnWhite(s string) DecoratedString {
	return DecoratedString(utils.ChangeBackgroundColor(s, ansi.WhiteBackground))
}
