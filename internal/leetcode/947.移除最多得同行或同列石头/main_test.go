package main

import "testing"

type args struct {
	stones [][]int
}

var tests = []struct {
	name    string
	args    args
	wantRes int
}{
	{"should be equal", args{[][]int{{0, 1}, {1, 0}, {1, 1}}}, 2},
	{"should be equal", args{[][]int{{0, 0}, {0, 1}, {1, 0}, {1, 2}, {2, 1}, {2, 2}}}, 5},
	{"should be equal", args{[][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}}, 3},
}

func Test_removeStonesFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := removeStonesFunc1(tt.args.stones); gotRes != tt.wantRes {
				t.Errorf("removeStonesFunc1() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
