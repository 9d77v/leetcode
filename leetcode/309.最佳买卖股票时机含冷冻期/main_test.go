package main

import (
	"testing"
)

type args struct {
	prices []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test", args{[]int{1, 2, 3, 0, 2}}, 3},
}

func Test_maxProfitFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitFunc1(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitFunc2(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
