package utils

import (
	"regexp"
)

func removeRegex(regex string, s string) string {
	re := regexp.MustCompile(regex)
	return re.ReplaceAllString(s, "")
}

// RemoveTextColors removes all text foreground coloring from a string
func RemoveTextColors(s string) string {
	return removeRegex(foregroundColorsRegex, s)
}

// RemoveBackgroundColors removes all background colors from a string
func RemoveBackgroundColors(s string) string {
	return removeRegex(backgroundColorsRegex, s)
}

// RemoveUnderlines removes all underlining from a string
func RemoveUnderlines(s string) string {
	return removeRegex(underlineRegex, s)
}

// RemoveBold removes all bolding from a string
func RemoveBold(s string) string {
	return removeRegex(boldRegex, s)
}

// RemoveInvert will remove all swapping of foreground and background colors from a string
func RemoveInvert(s string) string {
	return removeRegex(invertedRegex, s)
}
