package main

import (
	"testing"
)

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test f(1)", args{2}, 1},
	{"test f(5)", args{5}, 5},
	{"test f(45)", args{45}, 134903163},
	{"test f(95)", args{95}, 407059028},
}

func Test_fibFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibFunc2(tt.args.n); got != tt.want {
				t.Errorf("fibFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
