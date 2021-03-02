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
	want int
}{
	{"should be equal", args{"abcabcbb"}, 3},
	{"should be equal", args{"bbbbb"}, 1},
	{"should be equal", args{"pwwkew"}, 3},
	{"should be equal", args{"abba"}, 2},
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringFunc2(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
