package prism

import "testing"

func TestInRed(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want DecoratedString
	}{
		{"Red 'apple'", args{"apple"}, "\033[31mapple\033[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InRed(tt.args.s); got != tt.want {
				t.Errorf("\nInRed() = %v, want %v", got, tt.want)
			}
		})
	}
}
