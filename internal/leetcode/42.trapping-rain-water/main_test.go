package main

import "testing"

type args struct {
	height []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be ok", args{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}}, 6},
	{"should be ok", args{[]int{4, 2, 0, 3, 2, 5}}, 9},
}

func Test_trapFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trapFunc1(tt.args.height); got != tt.want {
				t.Errorf("trapFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
