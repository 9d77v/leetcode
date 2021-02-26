package main

import (
	"testing"
)

type args struct {
	matrix [][]int
	target int
}

var matrix1 = [][]int{
	{1, 4, 7, 11, 15},
	{2, 5, 8, 12, 19},
	{3, 6, 9, 16, 22},
	{10, 13, 14, 17, 24},
	{18, 21, 23, 26, 30},
}

var matrix2 = [][]int{{-5}}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"test target 5", args{matrix1, 5}, true},
	{"test target 20", args{matrix1, 20}, false},
	{"test target -5", args{matrix2, -5}, true},
}

func Test_findNumberIn2DArrayFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArrayFunc1(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArrayFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findNumberIn2DArrayFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumberIn2DArrayFunc2(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("findNumberIn2DArrayFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
