package ansi

import "testing"

func Test_castForegroundColorToAnsiCode(t *testing.T) {
	want := Code("31")
	if got := Code(RedText); got != want {
		t.Errorf("\napplyAnsiCode() = %v, want %v", got, want)
	}
}
