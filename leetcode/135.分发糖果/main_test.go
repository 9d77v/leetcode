package main

import (
	"testing"
)

type args struct {
	ratings []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test candy", args{nil}, 0},
	{"test candy", args{[]int{}}, 0},
	{"test candy", args{[]int{1}}, 1},
	{"test candy", args{[]int{1, 1}}, 2},
	{"test candy", args{[]int{1, 3}}, 3},
	{"test candy", args{[]int{1, 0, 2}}, 5},
	{"test candy", args{[]int{1, 2, 2}}, 4},
	{"test candy", args{[]int{1, 3, 4, 5, 2}}, 11},
	{"test candy", args{[]int{1, 3, 2, 2, 1}}, 7},
}

func Test_candyFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candyFunc1(tt.args.ratings); got != tt.want {
				t.Errorf("candyFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_candyFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := candyFunc2(tt.args.ratings); got != tt.want {
				t.Errorf("candyFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
