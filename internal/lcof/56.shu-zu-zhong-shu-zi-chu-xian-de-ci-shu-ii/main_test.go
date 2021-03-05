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
	{"should be equal", args{[]int{3, 4, 3, 3}}, 4},
	{"should be equal", args{[]int{9, 1, 7, 9, 7, 9, 7}}, 1},
}

func Test_singleNumber(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber(tt.args.nums); got != tt.want {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_singleNumberFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumberFunc2(tt.args.nums); got != tt.want {
				t.Errorf("singleNumberFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
