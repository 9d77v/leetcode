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
	want []int
}{
	{"should be equal", args{
		[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
	}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
	{"should be equal", args{
		[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
	}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
}

func Test_spiralOrderFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrderFunc1(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrderFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_spiralOrderFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := spiralOrderFunc2(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrderFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
