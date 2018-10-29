package prism

import (
	"testing"
)

func Test_complexStrings(t *testing.T) {
    t.Run("Underlined concatented with bold, all in red", func(t * testing.T) {
        got := ("Red with " + Underlined("underlined text") + " with "+ InBold("bold text")).InRed()
        want := DecoratedString("\033[31mRed with \033[4munderlined text\033[0m\033[31m with \033[1mbold text\033[0m")
        if got != want {
            t.Errorf("\napplyAnsiCode() = %v, \nwant %v\n", got, want)
        }
    })
}
