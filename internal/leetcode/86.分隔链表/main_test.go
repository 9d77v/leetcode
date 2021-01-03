package main

import (
	"reflect"
	"testing"

	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

type args struct {
	head *ListNode
	x    int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{NewList([]int{1, 4, 3, 2, 5, 2}), 3}, []int{1, 2, 2, 4, 3, 5}},
	{"should be equal", args{NewList([]int{1, 1}), 2}, []int{1, 1}},
	{"should be equal", args{NewList([]int{2, 1}), 2}, []int{1, 2}},
}

func Test_partition(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("partition() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
