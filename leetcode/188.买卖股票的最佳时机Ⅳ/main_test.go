package main

import (
	"testing"
)

type args struct {
	k      int
	prices []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test", args{2, []int{2, 4, 1}}, 2},
	{"test", args{2, []int{3, 2, 6, 5, 0, 3}}, 7},
}

func Test_maxProfitFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitFunc1(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitFunc2(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
