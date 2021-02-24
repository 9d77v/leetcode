package main

import (
	"reflect"
	"testing"
)

type args struct {
	matrix [][]int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"should be equal", args{[][]int{
		{2, 4, -1}, {-10, 5, 11}, {18, -7, 6},
	}}, [][]int{
		{2, -10, 18}, {4, 5, -7}, {-1, 11, 6},
	}},
	{"should be equal", args{[][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}}, [][]int{
		{1, 4, 7}, {2, 5, 8}, {3, 6, 9},
	}},
	{"should be equal", args{[][]int{
		{1, 2, 3}, {4, 5, 6},
	}}, [][]int{
		{1, 4}, {2, 5}, {3, 6},
	}},
}

func Test_transpose(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transpose(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("transpose() = %v, want %v", got, tt.want)
			}
		})
	}
}
