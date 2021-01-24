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
	{"should be equal", args{[]int{1, 3, 5, 4, 7}}, 3},
	{"should be equal", args{[]int{2, 2, 2, 2, 2}}, 1},
}

func Test_findLengthOfLCISFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCISFunc1(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCISFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLengthOfLCISFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLengthOfLCISFunc2(tt.args.nums); got != tt.want {
				t.Errorf("findLengthOfLCISFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
