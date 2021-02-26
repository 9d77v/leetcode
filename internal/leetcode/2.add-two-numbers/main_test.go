package main

import (
	"reflect"
	"testing"

	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{NewList([]int{2, 4, 3}), NewList([]int{5, 6, 4})}, []int{7, 0, 8}},
		{"should be equal", args{NewList([]int{0}), NewList([]int{0})}, []int{0}},
		{"should be equal", args{NewList([]int{9, 9, 9, 9, 9, 9, 9}), NewList([]int{9, 9, 9, 9})}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
