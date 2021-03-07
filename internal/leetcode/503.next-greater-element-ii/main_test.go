package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{[]int{1, 2, 1}}, []int{2, -1, 2}},
	{"should be equal", args{[]int{1, 2, 3, 4, 5}}, []int{2, 3, 4, 5, -1}},
}

func Test_nextGreaterElements(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElements(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextGreaterElementsFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextGreaterElementsFunc2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElementsFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
