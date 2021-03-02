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
	{"should be equal", args{1}, 1},
	{"should be equal", args{10}, 12},
}

func Test_nthUglyNumber(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumber(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nthUglyNumberFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nthUglyNumberFunc2(tt.args.n); got != tt.want {
				t.Errorf("nthUglyNumberFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
