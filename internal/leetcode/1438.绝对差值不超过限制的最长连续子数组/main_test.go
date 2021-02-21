package main

import (
	"testing"
)

type args struct {
	nums  []int
	limit int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{8, 2, 4, 7}, 4}, 2},
	{"should be equal", args{[]int{10, 1, 2, 4, 7, 2}, 5}, 4},
	{"should be equal", args{[]int{4, 2, 2, 2, 4, 4, 2, 2}, 0}, 3},
}

func Test_longestSubarray(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestSubarrayFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarrayFunc2(tt.args.nums, tt.args.limit); got != tt.want {
				t.Errorf("longestSubarrayFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
