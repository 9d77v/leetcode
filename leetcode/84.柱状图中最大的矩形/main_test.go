package main

import (
	"testing"
)

type args struct {
	heights []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test", args{[]int{2, 1, 5, 6, 2, 3}}, 10},
	{"test", args{[]int{2, 1, 2}}, 3},
	{"test", args{[]int{5, 4, 1, 2}}, 8},
	{"test", args{[]int{4, 2, 0, 3, 2, 4, 3, 4}}, 10},
	{"test", args{[]int{}}, 0},
}

func Test_largestRectangleAreaFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleAreaFunc1(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleAreaFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleAreaFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleAreaFunc2(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleAreaFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_largestRectangleAreaFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleAreaFunc3(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleAreaFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
