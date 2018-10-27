package utils

import (
	"github.com/juanrgon/prism/ansi"
)

func applyAnsiCode(code ansi.Code, text string) string {
	start := string(ansi.Begin + code + ansi.End)
	end := string(ansi.Reset)
	return start + text + end
}

// ChangeTextColor changes the text color of the string
func ChangeTextColor(s string, c ansi.ForegroundColor) string {
	changed := RemoveTextColors(s)
	return applyAnsiCode(ansi.Code(c), changed)
}

// ChangeBackgroundColor removes all background colors from a string
func ChangeBackgroundColor(s string, c ansi.BackgroundColor) string {
	changed := RemoveBackgroundColors(s)
	return applyAnsiCode(ansi.Code(c), changed)
}


// Underline returns the string underlined
func Underline(s string) string {
	changed := RemoveUnderlines(s)
	return applyAnsiCode(ansi.Code(ansi.Underlined), changed)
}

// Embolden returns the string in bold
func Embolden(s string) string {
	changed := RemoveBold(s)
	return applyAnsiCode(ansi.Code(ansi.Bold), changed)
}

// Invert will invert swapping of foreground and background colors from a string
func Invert(s string) string {
	changed := RemoveInvert(s)
	return applyAnsiCode(ansi.Code(ansi.Inverted), changed)
}
