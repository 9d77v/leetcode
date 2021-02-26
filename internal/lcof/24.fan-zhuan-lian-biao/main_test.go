package main

import (
	"reflect"
	"testing"

	. "github.com/9d77v/leetcode/pkg/algorithm/singlylinkedlist"
)

type args struct {
	head *ListNode
}

func Test_reverseListFunc1(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{NewList([]int{1, 2, 3, 4, 5})}, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListFunc1(tt.args.head); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("reverseListFunc1() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}

func Test_reverseListFunc2(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{NewList([]int{1, 2, 3, 4, 5})}, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListFunc2(tt.args.head); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("reverseListFunc2() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}

func Test_reverseListFunc3(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{NewList([]int{1, 2, 3, 4, 5})}, []int{5, 4, 3, 2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListFunc3(tt.args.head); !reflect.DeepEqual(got.ToArray(), tt.want) {
				t.Errorf("reverseListFunc3() = %v, want %v", got.ToArray(), tt.want)
			}
		})
	}
}
