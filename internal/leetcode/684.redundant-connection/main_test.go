package main

import (
	"reflect"
	"testing"
)

type args struct {
	edges [][]int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{[][]int{{1, 2}, {1, 3}, {2, 3}}}, []int{2, 3}},
	{"should be equal", args{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}}, []int{1, 4}},
	{"should be equal", args{[][]int{{9, 10}, {5, 8}, {2, 6}, {1, 5}, {3, 8}, {4, 9}, {8, 10}, {4, 10}, {6, 8}, {7, 9}}}, []int{4, 10}},
}

func Test_findRedundantConnectionFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantConnectionFunc1(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnectionFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRedundantConnectionFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRedundantConnectionFunc2(tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnectionFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
