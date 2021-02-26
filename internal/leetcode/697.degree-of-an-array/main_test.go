package main

import (
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{1, 2, 2, 3, 1}}, 2},
	{"should be equal", args{[]int{1, 2, 2, 3, 1, 4, 2}}, 6},
}

func Test_findShortestSubArray(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findShortestSubArray(tt.args.nums); got != tt.want {
				t.Errorf("findShortestSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
