package main

import (
	"testing"
)

type args struct {
	s        string
	wordDict []string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{"leetcode", []string{"leet", "code"}}, true},
	{"should be true", args{"applepenapple", []string{"apple", "pen"}}, true},
	{"should be false", args{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}}, false},
}

func Test_wordBreak(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreak(tt.args.s, tt.args.wordDict); got != tt.want {
				t.Errorf("wordBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
