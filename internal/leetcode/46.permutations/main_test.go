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
	{"should be equal", args{[]int{1, 2, 3}}, [][]int{
		{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 2, 1}, {3, 1, 2},
	}},
}

func Test_permute(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}