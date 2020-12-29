package main

import "testing"

type args struct {
	nums []int
	n    int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{[]int{1, 3}, 6}, 1},
	{"should be ok", args{[]int{1, 5, 10}, 20}, 2},
	{"should be ok", args{[]int{1, 2, 2, 5}, 5}, 0},
	{"should be ok", args{[]int{}, 8}, 4},
}

func Test_minPatchesFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minPatchesFunc1(tt.args.nums, tt.args.n); got != tt.want {
				t.Errorf("minPatchesFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
