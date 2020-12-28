package main

import (
	"testing"
)

type args struct {
	prices []int
	fee    int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test", args{[]int{1, 3, 2, 8, 4, 9}, 2}, 8},
}

func Test_maxProfitFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitFunc1(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfitFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxProfitFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitFunc2(tt.args.prices, tt.args.fee); got != tt.want {
				t.Errorf("maxProfitFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
