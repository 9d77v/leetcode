package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"should be equal", args{[]int{1, 1, 2}}, [][]int{
		{1, 1, 2}, {1, 2, 1}, {2, 1, 1},
	}},
	{"should be equal", args{[]int{1, 2, 3}}, [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1},
	}},
}

func Test_permuteUnique(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permuteUnique(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permuteUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
