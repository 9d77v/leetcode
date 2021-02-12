package main

import "testing"

type args struct {
	A []int
	K int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{1, 2, 1, 2, 3}, 2}, 7},
	{"should be equal", args{[]int{1, 2, 1, 3, 4}, 3}, 3},
}

func Test_subarraysWithKDistinct(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subarraysWithKDistinct(tt.args.A, tt.args.K); got != tt.want {
				t.Errorf("subarraysWithKDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
