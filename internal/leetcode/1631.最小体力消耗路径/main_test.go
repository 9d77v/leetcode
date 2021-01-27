package main

import (
	"testing"
)

type args struct {
	heights [][]int
}

var tests = []struct {
	name       string
	args       args
	wantResult int
}{
	{"should be equal", args{
		[][]int{
			{1, 2, 2}, {3, 8, 2}, {5, 3, 5},
		}}, 2},
	{"should be equal", args{
		[][]int{
			{1, 2, 3}, {3, 8, 4}, {5, 3, 5},
		}}, 1},
	{"should be equal", args{
		[][]int{
			{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1},
			{1, 2, 1, 2, 1}, {1, 1, 1, 2, 1},
		}}, 0},
	{"should be equal", args{
		[][]int{{1, 10, 6, 7, 9, 10, 4, 9}},
	}, 9},
	{"should be equal", args{
		[][]int{{4, 3, 4, 10, 5, 5, 9, 2}, {10, 8, 2, 10, 9, 7, 5, 6}, {5, 8, 10, 10, 10, 7, 4, 2}, {5, 1, 3, 1, 1, 3, 1, 9}, {6, 4, 10, 6, 10, 9, 4, 6}},
	}, 5},
}

func Test_minimumEffortPathFuc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPathFuc1(tt.args.heights); got != tt.wantResult {
				t.Errorf("minimumEffortPathFuc1() = %v, wantResult %v", got, tt.wantResult)
			}
		})
	}
}
