package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums [][]int
	r    int
	c    int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"should be equal", args{[][]int{{1, 2}, {3, 4}}, 1, 4}, [][]int{{1, 2, 3, 4}}},
	{"should be equal", args{[][]int{{1, 2}, {3, 4}}, 2, 4}, [][]int{{1, 2}, {3, 4}}},
}

func Test_matrixReshape(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixReshape(tt.args.nums, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixReshape() = %v, want %v", got, tt.want)
			}
		})
	}
}
