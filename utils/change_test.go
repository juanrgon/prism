package utils

import (
	"testing"

	"github.com/juanrgon/prism/ansi"
)

func TestChangeTextColor(t *testing.T) {
	type args struct {
		s string
		c ansi.ForegroundColor
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Change uncolored string to red", args{"apple", ansi.RedText}, "\033[31mapple"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChangeTextColor(tt.args.s, tt.args.c); got != tt.want {
				t.Errorf("\nChangeTextColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_castForegroundColorToAnsiCode(t *testing.T) {
	want := ansi.Code("31")
	if got := ansi.Code(ansi.RedText); got != want {
		t.Errorf("\napplyAnsiCode() = %v, want %v", got, want)
	}
}

func Test_applyAnsiCode(t *testing.T) {
	type args struct {
		code ansi.Code
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Apply red text ansi code to 'apple'", args{ansi.Code(ansi.RedText), "apple"}, "\033[31mapple\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyAnsiCode(tt.args.code, tt.args.text); got != tt.want {
				t.Errorf("\napplyAnsiCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
