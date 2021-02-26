package main

import (
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[]int{1, 1, 0, 1, 1, 1}}, 3},
}

func Test_findMaxConsecutiveOnes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnes(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findMaxConsecutiveOnesFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxConsecutiveOnesFunc2(tt.args.nums); got != tt.want {
				t.Errorf("findMaxConsecutiveOnesFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
