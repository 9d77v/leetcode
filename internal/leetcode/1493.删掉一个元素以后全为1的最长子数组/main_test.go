package main

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{1, 1, 0, 1}}, 3},
	{"should be equal", args{[]int{0, 1, 1, 1, 0, 1, 1, 0, 1}}, 5},
	{"should be equal", args{[]int{1, 1, 1}}, 2},
	{"should be equal", args{[]int{1, 1, 0, 0, 1, 1, 1, 0, 1}}, 4},
	{"should be equal", args{[]int{0, 0, 0}}, 0},
}

func Test_longestSubarray(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
