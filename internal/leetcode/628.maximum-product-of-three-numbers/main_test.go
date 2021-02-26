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
	{"should be ok", args{[]int{1, 2, 3}}, 6},
	{"should be ok", args{[]int{1, 2, 3, 4}}, 24},
	{"should be ok", args{[]int{-100, -98, -1, 2, 3, 4}}, 39200},
	{"should be ok", args{[]int{0 - 8, -7, -2, 10, 20}}, 1120},
	{"should be ok", args{[]int{-2, -3, -4}}, -24},
}

func Test_maximumProduct(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProduct(tt.args.nums); got != tt.want {
				t.Errorf("maximumProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumProductFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProductFunc2(tt.args.nums); got != tt.want {
				t.Errorf("maximumProductFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maximumProductFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumProductFunc3(tt.args.nums); got != tt.want {
				t.Errorf("maximumProductFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
