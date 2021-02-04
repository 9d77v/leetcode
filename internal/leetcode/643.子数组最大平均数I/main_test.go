package main

import "testing"

type args struct {
	nums []int
	k    int
}

var tests = []struct {
	name string
	args args
	want float64
}{
	{"should be equal", args{[]int{1, 12, -5, -6, 50, 3}, 4}, 12.75},
	{"should be equal", args{[]int{-1}, 1}, -1.0},
}

func Test_findMaxAverage(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
