package main

import (
	"reflect"
	"testing"
)

func Test_sortItems(t *testing.T) {
	type args struct {
		n           int
		m           int
		group       []int
		beforeItems [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"should be equal",
			args{8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3, 6}, {}, {}, {}}},
			[]int{6, 3, 4, 5, 2, 0, 7, 1}},
		{"should be equal",
			args{8, 2, []int{-1, -1, 1, 0, 0, 1, 0, -1}, [][]int{{}, {6}, {5}, {6}, {3}, {}, {4}, {}}},
			[]int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortItems(tt.args.n, tt.args.m, tt.args.group, tt.args.beforeItems); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
