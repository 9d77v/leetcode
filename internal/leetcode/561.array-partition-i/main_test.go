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
	{"should be equal", args{[]int{1, 4, 3, 2}}, 4},
	{"should be equal", args{[]int{6, 2, 6, 5, 1, 2}}, 9},
}

func Test_arrayPairSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum(tt.args.nums); got != tt.want {
				t.Errorf("arrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
