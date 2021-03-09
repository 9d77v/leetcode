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
	{"should be equal", args{"       "}, ""},
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

func Test_reverseWordsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWordsFunc2(tt.args.s); got != tt.want {
				t.Errorf("reverseWordsFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseWordsFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWordsFunc3(tt.args.s); got != tt.want {
				t.Errorf("reverseWordsFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
