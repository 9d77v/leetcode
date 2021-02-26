package main

import "testing"

type args struct {
	s     string
	pairs [][]int
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"should be equal", args{"dcab", [][]int{{0, 3}, {1, 2}}}, "bacd"},
	{"should be equal", args{"dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}}, "abcd"},
	{"should be equal", args{"cba", [][]int{{0, 1}, {1, 2}}}, "abc"},
}

func Test_smallestStringWithSwaps(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestStringWithSwaps(tt.args.s, tt.args.pairs); got != tt.want {
				t.Errorf("smallestStringWithSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
