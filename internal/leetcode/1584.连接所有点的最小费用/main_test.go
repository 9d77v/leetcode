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
	{"should be equal", args{[][]int{
		{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0},
	}}, 20},
	{"should be equal", args{[][]int{
		{3, 12}, {-2, 5}, {-4, 1},
	}}, 18},
	{"should be equal", args{[][]int{
		{0, 0}, {1, 1}, {1, 0}, {-1, 1},
	}}, 4},
	{"should be equal", args{[][]int{
		{-1000000, -1000000}, {1000000, 1000000},
	}}, 4000000},
	{"should be equal", args{[][]int{
		{0, 0},
	}}, 0},
}

func Test_minCostConnectPointsFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPointsFunc1(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPointsFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minCostConnectPointsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostConnectPointsFunc2(tt.args.points); got != tt.want {
				t.Errorf("minCostConnectPointsFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
