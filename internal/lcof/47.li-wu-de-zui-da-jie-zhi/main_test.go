package main

import (
	"testing"
)

type args struct {
	grid [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}}, 12},
	{"should be equal", args{[][]int{
		{1, 2, 5}, {3, 2, 1},
	}}, 9},
}

func Test_maxValue(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValue(tt.args.grid); got != tt.want {
				t.Errorf("maxValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxValueFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxValueFunc2(tt.args.grid); got != tt.want {
				t.Errorf("maxValueFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
