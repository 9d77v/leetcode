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
	{"should be ok", args{"abcabcbb"}, 3},
	{"should be ok", args{"bbbbb"}, 1},
	{"should be ok", args{"pwwkew"}, 3},
	{"should be ok", args{""}, 0},
	{"should be ok", args{" "}, 1},
	{"should be ok", args{"dvdf"}, 3},
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
