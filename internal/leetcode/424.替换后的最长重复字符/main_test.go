package main

import (
	"testing"
)

type args struct {
	s string
	k int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{"ABAB", 2}, 4},
	{"should be equal", args{"AABABBA", 1}, 4},
}

func Test_characterReplacement(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := characterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
