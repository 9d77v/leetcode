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
	{"should be true", args{[]int{1, 2, 3, 1}}, true},
	{"should be false", args{[]int{1, 2, 3, 4}}, false},
	{"should be true", args{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}}, true},
}

func Test_containsDuplicateFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicateFunc1(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicateFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_containsDuplicateFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicateFunc2(tt.args.nums); got != tt.want {
				t.Errorf("containsDuplicateFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
