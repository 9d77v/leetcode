package main

import (
	"reflect"
	"testing"
)

type args struct {
	n     int
	edges [][]int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"should be equal", args{5, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 2}, {0, 3, 2}, {0, 4, 3}, {3, 4, 3}, {1, 4, 6}}},
		[][]int{{0, 1}, {2, 3, 4, 5}}},
	{"should be equal", args{4, [][]int{{0, 1, 1}, {1, 2, 1}, {2, 3, 1}, {0, 3, 1}}},
		[][]int{{}, {0, 1, 2, 3}}},
	{"should be equal", args{2, [][]int{{0, 1, 3}}},
		[][]int{{0}, {}}},
	{"should be equal", args{4, [][]int{{0, 1, 1}, {0, 2, 2}, {0, 3, 3}}}, [][]int{{0, 1, 2}, {}}},
}

func Test_findCriticalAndPseudoCriticalEdgesFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCriticalAndPseudoCriticalEdgesFunc1(tt.args.n, tt.args.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCriticalAndPseudoCriticalEdgesFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
