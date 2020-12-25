package main

import "testing"

type args struct {
	g []int
	s []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"test1", args{[]int{1, 2, 3}, []int{1, 1}}, 1},
	{"test2", args{[]int{1, 2}, []int{1, 2, 3}}, 2},
	{"test3", args{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}}, 2},
}

func Test_findContentChildrenFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContentChildrenFunc1(tt.args.g, tt.args.s); got != tt.want {
				t.Errorf("findContentChildrenFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
