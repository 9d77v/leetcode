package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums1 []int
	nums2 []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be ok", args{[]int{4, 1, 2}, []int{1, 3, 4, 2}}, []int{-1, 3, -1}},
	{"should be ok", args{[]int{2, 4}, []int{1, 2, 3, 4}}, []int{3, -1}},
}

func Test_nextGreaterElement(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElement(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
