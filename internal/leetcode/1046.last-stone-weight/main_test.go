package main

import (
	"testing"
)

type args struct {
	stones []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{[]int{2, 7, 4, 1, 8, 1}}, 1},
}

func Test_lastStoneWeightFun1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeightFun1(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeightFun1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lastStoneWeightFun2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeightFun2(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeightFun2() = %v, want %v", got, tt.want)
			}
		})
	}
}
