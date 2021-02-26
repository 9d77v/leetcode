package main

import (
	"testing"
)

type args struct {
	num uint32
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{0b00000000000000000000000000001011}, 3},
	{"should be ok", args{0b00000000000000000000000010000000}, 1},
	{"should be ok", args{0b11111111111111111111111111111101}, 31},
}

func Test_hammingWeight(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeightFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeightFunc2(tt.args.num); got != tt.want {
				t.Errorf("hammingWeightFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
