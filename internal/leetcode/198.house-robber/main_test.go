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
	{"should be equal", args{[]int{1, 2, 3, 1}}, 4},
	{"should be equal", args{[]int{2, 7, 9, 3, 1}}, 12},
	{"should be equal", args{[]int{2, 1, 1, 2}}, 4},
}

func Test_rob(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_robFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robFunc2(tt.args.nums); got != tt.want {
				t.Errorf("robFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
