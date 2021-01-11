package main

import (
	"reflect"
	"testing"

	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

type args struct {
	l1 *ListNode
	l2 *ListNode
}

func Test_mergeTwoListsFunc1(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{NewList([]int{1, 2, 3}), NewList([]int{1, 3, 4})}, []int{1, 1, 2, 3, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoListsFunc1(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("mergeTwoListsFunc1() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}

func Test_mergeTwoListsFunc2(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{NewList([]int{1, 2, 3}), NewList([]int{1, 3, 4})}, []int{1, 1, 2, 3, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoListsFunc2(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("mergeTwoListsFunc2() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
