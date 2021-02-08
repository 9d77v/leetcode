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
	want bool
}{
	{"should be true", args{[]int{4, 2, 3}}, true},
	{"should be false", args{[]int{4, 2, 1}}, false},
	{"should be true", args{[]int{5, 7, 1, 8}}, true},
	{"should be false", args{[]int{3, 4, 2, 3}}, false},
}

func Test_checkPossibility(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPossibility(tt.args.nums); got != tt.want {
				t.Errorf("checkPossibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
