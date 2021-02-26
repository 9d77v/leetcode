package main

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{1, 7, 3, 6, 5, 6}}, 3},
	{"should be equal", args{[]int{1, 2, 3}}, -1},
	{"should be equal", args{[]int{-1, -1, -1, -1, -1, 0}}, 2},
}

func Test_pivotIndex(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
