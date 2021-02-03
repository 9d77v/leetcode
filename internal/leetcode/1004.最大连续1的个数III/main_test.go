package main

import (
	"testing"
)

type args struct {
	A []int
	K int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2}, 6},
	{"should be ok", args{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3}, 10},
}

func Test_longestOnes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestOnes(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("longestOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
