package main

import (
	"testing"
)

type args struct {
	s string
	n int
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"should be equal", args{"abcdefg", 2}, "cdefgab"},
	{"should be equal", args{"lrloseumgh", 6}, "umghlrlose"},
}

func Test_reverseLeftWords(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWords(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseLeftWordsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseLeftWordsFunc2(tt.args.s, tt.args.n); got != tt.want {
				t.Errorf("reverseLeftWordsFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
