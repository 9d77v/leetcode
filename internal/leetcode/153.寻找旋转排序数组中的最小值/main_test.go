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
	{"should be ok", args{[]int{3, 4, 5, 1, 2}}, 1},
	{"should be ok", args{[]int{4, 5, 6, 7, 0, 1, 2}}, 0},
	{"should be ok", args{[]int{1}}, 1},
}

func Test_findMin(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
