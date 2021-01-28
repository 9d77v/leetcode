package main

import (
	"testing"
)

type args struct {
	text1 string
	text2 string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{"abcde", "ace"}, 3},
	{"should be equal", args{"abc", "abc"}, 3},
	{"should be equal", args{"abc", "def"}, 0},
}

func Test_longestCommonSubsequenceFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequenceFunc1(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestCommonSubsequenceFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonSubsequenceFunc2(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("longestCommonSubsequenceFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
