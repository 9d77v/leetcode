package main

import (
	"testing"
)

func Test_minimumHammingDistance(t *testing.T) {
	type args struct {
		source       []int
		target       []int
		allowedSwaps [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"should be equal", args{[]int{1, 2, 3, 4}, []int{2, 1, 4, 5}, [][]int{{0, 1}, {2, 3}}}, 1},
		{"should be equal", args{[]int{1, 2, 3, 4}, []int{1, 3, 2, 4}, [][]int{}}, 2},
		{"should be equal", args{[]int{5, 1, 2, 4, 3}, []int{1, 5, 4, 2, 3}, [][]int{{0, 4}, {4, 2}, {1, 3}, {1, 4}}}, 0},
		{"should be equal", args{[]int{2, 3, 1}, []int{1, 2, 2}, [][]int{{0, 2}, {1, 2}}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumHammingDistance(tt.args.source, tt.args.target, tt.args.allowedSwaps); got != tt.want {
				t.Errorf("minimumHammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
