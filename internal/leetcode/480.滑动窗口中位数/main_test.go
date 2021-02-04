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
	want []float64
}{
	{"should be ok", args{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3}, []float64{1, -1, -1, 3, 5, 6}},
}

func Test_medianSlidingWindow(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindow(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_medianSlidingWindowFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindowFunc2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindowFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_medianSlidingWindowFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := medianSlidingWindowFunc3(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindowFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
