package main

import (
	"reflect"
	"testing"

	. "github.com/9d77v/leetcode/pkg/algorithm/complexlist"
)

func Test_copyRandomList(t *testing.T) {
	a := 3
	b := 0
	l := NewComplexList([][]*int{
		{&a, nil}, {&a, &b}, {&a, nil},
	})
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{"should be equal", args{l}, l},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomList() = %v, want %v", got, tt.want)
			}
		})
	}
}
