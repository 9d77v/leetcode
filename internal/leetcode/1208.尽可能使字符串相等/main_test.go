package main

import "testing"

type args struct {
	s       string
	t       string
	maxCost int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{"abcd", "bcdf", 3}, 3},
	{"should be equal", args{"abcd", "cdef", 3}, 1},
	{"should be equal", args{"abcd", "acde", 0}, 1},
}

func Test_equalSubstring(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := equalSubstring(tt.args.s, tt.args.t, tt.args.maxCost); got != tt.want {
				t.Errorf("equalSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
