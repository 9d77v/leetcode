package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
	k    int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be ok", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
	{"should be ok", args{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
}

func Test_rotateFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateFunc1(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("rotate failed,got:%v,want:%v", tt.args.nums, tt.want)
			}
		})
	}
}

func Test_rotateFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateFunc2(tt.args.nums, tt.args.k)
		})
	}
}

func Test_rotateFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateFunc3(tt.args.nums, tt.args.k)
		})
	}
}
