package main

import "testing"

type args struct {
	flowerbed []int
	n         int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be ok", args{[]int{1, 0, 0, 0, 1}, 1}, true},
	{"should be ok", args{[]int{1, 0, 0, 0, 1}, 2}, false},
	{"should be ok", args{[]int{1, 0, 1, 0, 1, 0, 1}, 0}, true},
	{"should be ok", args{[]int{1, 0, 0, 0, 1, 0, 0}, 2}, true},
}

func Test_canPlaceFlowersFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowersFunc1(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowersFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
