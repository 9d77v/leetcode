package main

import "testing"

type args struct {
	s string
	t string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"should be equal", args{"ADOBECODEBANC", "ABC"}, "BANC"},
	{"should be equal", args{"a", "a"}, "a"},
	{"should be equal", args{"ab", "b"}, "b"},
	{"should be equal", args{"aa", "aa"}, "aa"},
	{"should be equal", args{"bba", "ab"}, "ba"},
	{"should be equal", args{"acbbaca", "aba"}, "baca"},
}

func Test_minWindowFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindowFunc1(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindowFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
