package main

import (
	"reflect"
	"testing"

	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

type args struct {
	head *ListNode
	val  int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{NewList([]int{4, 5, 1, 9}), 5}, []int{4, 1, 9}},
	{"should be equal", args{NewList([]int{4, 5, 1, 9}), 1}, []int{4, 5, 9}},
}

func Test_deleteNode(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteNode(tt.args.head, tt.args.val); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("deleteNode() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
