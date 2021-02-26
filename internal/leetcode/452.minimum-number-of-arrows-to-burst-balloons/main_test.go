package main

import (
	"testing"
)

type args struct {
	points [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{[][]int{
		{10, 16}, {2, 8}, {1, 6}, {7, 12},
	}}, 2},
	{"should be ok", args{[][]int{
		{1, 2}, {3, 4}, {5, 6}, {7, 8},
	}}, 4},
	{"should be ok", args{[][]int{
		{1, 2}, {2, 3}, {3, 4}, {4, 5},
	}}, 2},
	{"should be ok", args{[][]int{
		{1, 2},
	}}, 1},
	{"should be ok", args{[][]int{
		{2, 3}, {2, 3},
	}}, 1},
}

func Test_findMinArrowShots(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMinArrowShots(tt.args.points); got != tt.want {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tt.want)
			}
		})
	}
}
