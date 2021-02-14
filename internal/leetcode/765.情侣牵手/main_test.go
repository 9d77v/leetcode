package main

import (
	"testing"
)

type args struct {
	row []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{0, 2, 1, 3}}, 1},
	{"should be equal", args{[]int{3, 2, 0, 1}}, 0},
	{"should be equal", args{[]int{5, 4, 2, 6, 3, 1, 0, 7}}, 2},
	{"should be equal", args{[]int{1, 4, 0, 5, 8, 7, 6, 3, 2, 9}}, 3},
}

func Test_minSwapsCouples(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwapsCouples(tt.args.row); got != tt.want {
				t.Errorf("minSwapsCouples() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSwapsCouplesFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSwapsCouplesFunc2(tt.args.row); got != tt.want {
				t.Errorf("minSwapsCouplesFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
