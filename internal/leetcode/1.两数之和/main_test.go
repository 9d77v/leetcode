package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums   []int
	target int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
	{"should be equal", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
	{"should be equal", args{[]int{3, 3}, 6}, []int{0, 1}},
}

func Test_twoSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_twoSumFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumFunc2(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
