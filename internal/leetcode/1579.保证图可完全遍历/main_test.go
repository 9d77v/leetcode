package main

import (
	"testing"
)

type args struct {
	n     int
	edges [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{
		4, [][]int{
			{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4},
		},
	}, 2},
	{"should be equal", args{
		4, [][]int{
			{3, 1, 2}, {3, 2, 3}, {1, 1, 4}, {2, 1, 4},
		},
	}, 0},
	{"should be equal", args{
		4, [][]int{
			{3, 2, 2}, {1, 1, 2}, {2, 3, 4},
		},
	}, -1},
}

func Test_maxNumEdgesToRemove(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumEdgesToRemove(tt.args.n, tt.args.edges); got != tt.want {
				t.Errorf("maxNumEdgesToRemove() = %v, want %v", got, tt.want)
			}
		})
	}
}
