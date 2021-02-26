package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
	k    int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be ok", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []int{3, 3, 5, 5, 6, 7}},
	{"should be ok", args{[]int{1}, 1}, []int{1}},
	{"should be ok", args{[]int{1, -1}, 1}, []int{1, -1}},
	{"should be ok", args{[]int{9, 11}, 2}, []int{11}},
	{"should be ok", args{[]int{4, -2}, 2}, []int{4}},
}

func Test_maxSlidingWindowFunc0(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindowFunc0(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindowFunc0() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindowFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindowFunc1(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindowFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindowFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindowFunc2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindowFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxSlidingWindowFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSlidingWindowFunc3(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindowFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
