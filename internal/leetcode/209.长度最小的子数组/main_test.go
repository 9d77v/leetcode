package main

import (
	"testing"
)

type args struct {
	s    int
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{7, []int{2, 3, 1, 2, 4, 3}}, 2},
	{"should be equal", args{77, []int{2, 3, 1, 2, 4, 3}}, 0},
	{"should be equal", args{4, []int{1, 4, 4}}, 1},
	{"should be equal", args{11, []int{1, 2, 3, 4, 5}}, 3},
}

func Test_minSubArrayLenFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLenFunc1(tt.args.s, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLenFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
