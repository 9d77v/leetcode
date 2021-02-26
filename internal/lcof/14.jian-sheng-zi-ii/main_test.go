package main

import "testing"

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{2}, 1},
	{"should be ok", args{10}, 36},
}

func Test_cuttingRopeFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cuttingRopeFunc1(tt.args.n); got != tt.want {
				t.Errorf("cuttingRopeFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
