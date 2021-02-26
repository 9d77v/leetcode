package main

import "testing"

type args struct {
	s1 string
	s2 string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{"ab", "eidbaooo"}, true},
	{"should be false", args{"ab", "eidboaoo"}, false},
}

func Test_checkInclusion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusion(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
