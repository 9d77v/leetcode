package main

import (
	"testing"
)

type args struct {
	numCourses    int
	prerequisites [][]int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{2, [][]int{{1, 0}}}, true},
	{"should be false", args{2, [][]int{{1, 0}, {0, 1}}}, false},
}

func Test_canFinishFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinishFunc1(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinishFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_canFinishFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinishFunc2(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinishFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
