package main

import "testing"

type args struct {
	nums   []int
	target int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{5, 7, 7, 8, 8, 10}, 8}, 2},
	{"should be equal", args{[]int{5, 7, 7, 8, 8, 10}, 6}, 0},
}

func Test_search(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
