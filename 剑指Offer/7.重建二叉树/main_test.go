package main

import (
	"reflect"
	"testing"
)

type args struct {
	preorder []int
	inorder  []int
}

var tests = []struct {
	name string
	args args
	want []interface{}
}{
	{"test", args{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}}, []interface{}{3, 9, 20, nil, nil, 15, 7}},
	{"test", args{[]int{3, 9, 20, 15}, []int{9, 3, 15, 20}}, []interface{}{3, 9, 20, nil, nil, 15}},
}

func Test_buildTreeFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTreeFunc1(tt.args.preorder, tt.args.inorder)
			if !reflect.DeepEqual(got.BFSWithNil(), tt.want) {
				t.Errorf("buildTreeFunc1() = %v, want %v", got.BFSWithNil(), tt.want)
			}
		})
	}
}

func Test_buildTreeFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := buildTreeFunc2(tt.args.preorder, tt.args.inorder)
			if !reflect.DeepEqual(got.BFSWithNil(), tt.want) {
				t.Errorf("buildTreeFunc2() = %v, want %v", got.BFSWithNil(), tt.want)
			}
		})
	}
}
