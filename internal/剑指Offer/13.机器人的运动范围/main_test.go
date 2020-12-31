package main

import "testing"

type args struct {
	m int
	n int
	k int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{2, 3, 1}, 3},
	{"should be ok", args{3, 1, 0}, 1},
}

func Test_movingCount(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := movingCount(tt.args.m, tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("movingCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
