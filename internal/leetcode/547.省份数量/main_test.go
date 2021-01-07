package main

import (
	"testing"
)

type args struct {
	isConnected [][]int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"should be equal", args{[][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}}, 2},
	{"should be equal", args{[][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}}, 3},
}

func Test_findCircleNumFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNumFunc1(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNumFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCircleNumFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNumFunc2(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNumFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findCircleNumFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findCircleNumFunc3(tt.args.isConnected); got != tt.want {
				t.Errorf("findCircleNumFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
