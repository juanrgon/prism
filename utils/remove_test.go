package utils

import (
	"testing"
)

func Test_removeRegex(t *testing.T) {
	type args struct {
		regex string
		s     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Remove 'apple'", args{"apple", "an apple tree"}, "an  tree"},
		{"Remove escape char'", args{"\033", "aa\033aa"}, "aaaa"},
		{"Remove text ansi escape code'", args{foregroundColorsRegex, "\033[35mred text"}, "red text"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeRegex(tt.args.regex, tt.args.s); got != tt.want {
				t.Errorf("removeRegex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveTextColors(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Remove from uncolored string", args{"some string"}, "some string"},
		{"Remove from underlined string", args{"\033[4munderlined string"}, "\033[4munderlined string"},
		{"Remove from colored string", args{"\033[31mred text"}, "red text"},
		{"Remove from multicolored string'", args{"\033[31mred text \033[32mgreen text"}, "red text green text"},
		{"Remove from multicolored underlined string'", args{"\033[4m\033[31mred text \033[32mgreen text"}, "\033[4mred text green text"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveTextColors(tt.args.s); got != tt.want {
				t.Errorf("\nRemoveTextColors() = '%v', 'want %v'", got, tt.want)
			}
		})
	}
}

func TestRemoveBackgroundColors(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Remove from uncolored string", args{"some string"}, "some string"},
		{"Remove from underlined string", args{"\033[4munderlined string"}, "\033[4munderlined string"},
		{"Remove from colored string", args{"\033[31mred text"}, "\033[31mred text"},
		{"Remove from string with colored background'", args{"\033[41mred text \033[42mgreen text"}, "red text green text"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveBackgroundColors(tt.args.s); got != tt.want {
				t.Errorf("\nRemoveBackgroundColors() = %v, want %v", got, tt.want)
			}
		})
	}
}
