package main

import (
	"testing"
)

type args struct {
	strs []string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]string{"tars", "rats", "arts", "star"}}, 2},
	{"should be equal", args{[]string{"omv", "ovm"}}, 1},
}

func Test_numSimilarGroups(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSimilarGroups(tt.args.strs); got != tt.want {
				t.Errorf("numSimilarGroups() = %v, want %v", got, tt.want)
			}
		})
	}
}
