package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums1 []int
	m     int
	nums2 []int
	n     int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"test ok", args{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3}, []int{1, 2, 2, 3, 5, 6}},
}

func Test_merge(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if !reflect.DeepEqual(tt.args.nums1, tt.want) {
				t.Errorf("two array should equal,got:%v+ want:%v+", tt.args.nums1, tt.want)
			}
		})
	}
}
