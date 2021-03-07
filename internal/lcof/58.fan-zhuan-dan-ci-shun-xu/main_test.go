package main

import (
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"should be equal", args{"the sky is blue"}, "blue is sky the"},
	{"should be equal", args{"  hello world!  "}, "world! hello"},
	{"should be equal", args{"a good   example"}, "example good a"},
}

func Test_reverseWords(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
