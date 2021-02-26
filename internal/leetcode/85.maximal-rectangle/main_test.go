package main

import "testing"

type args struct {
	matrix [][]byte
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test", args{[][]byte{
		{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'},
	}}, 6},
	{"test", args{[][]byte{}}, 0},
	{"test", args{[][]byte{{'1'}}}, 1},
	{"test", args{[][]byte{
		{'1', '0', '1', '1', '1'}, {'0', '1', '0', '1', '0'}, {'1', '1', '0', '1', '1'}, {'1', '1', '0', '1', '1'},
		{'0', '1', '1', '1', '1'},
	}}, 6},
}

func Test_maximalRectangleFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximalRectangleFunc1(tt.args.matrix); got != tt.want {
				t.Errorf("maximalRectangleFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
