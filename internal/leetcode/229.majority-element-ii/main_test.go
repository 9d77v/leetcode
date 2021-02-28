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
	{"should be equal", args{[]int{3, 2, 3}}, []int{3}},
	{"should be equal", args{[]int{1}}, []int{1}},
	{"should be equal", args{[]int{1, 1, 1, 3, 3, 2, 2, 2}}, []int{1, 2}},
	{"should be equal", args{[]int{1, 1, 2, 2, 7, 7, 8, 8, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3}}, []int{3, 9}},
}

func Test_majorityElement(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
