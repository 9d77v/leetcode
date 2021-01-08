package main

import (
	"reflect"
	"testing"

	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

type args struct {
	head *ListNode
	k    int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{NewList([]int{1, 2, 3, 4, 5}), 2}, []int{4, 5}},
}

func Test_getKthFromEndFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEndFunc1(tt.args.head, tt.args.k); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("getKthFromEndFunc1() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}

func Test_getKthFromEndFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getKthFromEndFunc2(tt.args.head, tt.args.k); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("getKthFromEndFunc2() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
