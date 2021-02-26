package main

import "testing"

type args struct {
	s string
	p string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be false", args{"aa", "a"}, false},
	{"should be true", args{"aa", "a*"}, true},
	{"should be true", args{"ab", ".*"}, true},
	{"should be true", args{"aab", "c*a*b"}, true},
	{"should be false", args{"mississippi", "mis*is*p*."}, false},
}

func Test_isMatch(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.args.s, tt.args.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
