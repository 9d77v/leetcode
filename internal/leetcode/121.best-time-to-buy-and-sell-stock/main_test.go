package main

import (
	"testing"
)

type args struct {
	prices []int
}

var tests = []struct {
	name    string
	args    args
	wantMax int
}{
	{"test", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
	{"test", args{[]int{7, 1, 5, 3, 6, 4}}, 5},
}

func Test_maxProfitFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxProfitFunc1(tt.args.prices); gotMax != tt.wantMax {
				t.Errorf("maxProfitFunc1() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}

func Test_maxProfitFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotMax := maxProfitFunc2(tt.args.prices); gotMax != tt.wantMax {
				t.Errorf("maxProfitFunc2() = %v, want %v", gotMax, tt.wantMax)
			}
		})
	}
}
