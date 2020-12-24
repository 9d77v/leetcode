package main

import (
	"reflect"
	"testing"

	"github.com/9d77v/leetcode/lib/singlylinkedlist"
)

var l *singlylinkedlist.ListNode

func TestMain(m *testing.M) {
	l = singlylinkedlist.NewListNode(1)
	l.Next = singlylinkedlist.NewListNode(3)
	l.Next.Next = singlylinkedlist.NewListNode(2)

	m.Run()
}

var tests = []struct {
	name  string
	input []int
	want  []int
}{
	{"test reverse print", []int{1, 3, 2}, []int{2, 3, 1}},
	{"test reverse print", []int{8, 4, 0, 6, 5, 6, 5, 7}, []int{7, 5, 6, 5, 6, 0, 4, 8}},
}

func Test_reversePrint(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrintFunc1(getList(tt.input)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePrintFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrintFunc2(getList(tt.input)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrintFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reversePrintFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reversePrintFunc3(getList(tt.input)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reversePrintFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getList(input []int) *singlylinkedlist.ListNode {
	n := &singlylinkedlist.ListNode{}
	m := n
	for i, num := range input {
		n.Val = num
		if i != len(input)-1 {
			n.Next = &singlylinkedlist.ListNode{}
			n = n.Next
		}
	}
	return m
}
