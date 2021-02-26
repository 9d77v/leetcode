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
	{"test findRepeatNumber", args{[]int{2, 3, 5, 4, 3, 2, 6, 7}}, 3},
	{"test findRepeatNumber", args{[]int{1, 2, 2}}, 2},
}

func Test_findRepeatNumberFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberFunc1(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumberFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRepeatNumberFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatNumberFunc2(tt.args.nums); got != tt.want {
				t.Errorf("findRepeatNumberFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
