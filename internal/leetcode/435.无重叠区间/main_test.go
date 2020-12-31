package main

import (
	"testing"
)

type args struct {
	intervals [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{[][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	}}, 1},
	{"should be ok", args{[][]int{
		{1, 2}, {1, 2}, {1, 2},
	}}, 2},
	{"should be ok", args{[][]int{
		{1, 2}, {2, 3},
	}}, 0},
}

func Test_eraseOverlapIntervalsFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervalsFunc1(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervalsFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_eraseOverlapIntervalsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := eraseOverlapIntervalsFunc2(tt.args.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervalsFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
