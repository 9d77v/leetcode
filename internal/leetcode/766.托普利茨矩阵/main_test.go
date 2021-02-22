package main

import (
	"testing"
)

type args struct {
	matrix [][]int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{[][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}}}, true},
	{"should be false", args{[][]int{{1, 2}, {2, 2}}}, false},
	{"should be true", args{[][]int{{18}, {66}}}, true},
	{"should be false", args{[][]int{{11, 74, 0, 93}, {40, 11, 74, 7}}}, false},
}

func Test_isToeplitzMatrix(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isToeplitzMatrixFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isToeplitzMatrixFunc2(tt.args.matrix); got != tt.want {
				t.Errorf("isToeplitzMatrixFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
