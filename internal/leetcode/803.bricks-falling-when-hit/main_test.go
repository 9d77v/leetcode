package main

import (
	"reflect"
	"testing"
)

type args struct {
	grid [][]int
	hits [][]int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{
		[][]int{{1, 0, 0, 0}, {1, 1, 1, 0}},
		[][]int{{1, 0}},
	}, []int{2}},
	{"should be equal", args{
		[][]int{{1, 0, 0, 0}, {1, 1, 0, 0}},
		[][]int{{1, 1}, {1, 0}},
	}, []int{0, 0}},
}

func Test_hitBricks(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hitBricks(tt.args.grid, tt.args.hits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hitBricks() = %v, want %v", got, tt.want)
			}
		})
	}
}
