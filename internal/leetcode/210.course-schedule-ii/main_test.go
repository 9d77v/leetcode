package main

import (
	"reflect"
	"testing"
)

type args struct {
	numCourses    int
	prerequisites [][]int
}

var tests = []struct {
	name  string
	args  args
	want  []int
	want2 []int
}{
	{"should be equal", args{2, [][]int{{1, 0}}}, []int{0, 1}, []int{0, 1}},
	{"should be equal", args{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}}, []int{0, 1, 2, 3}, []int{0, 2, 1, 3}},
}

func Test_findOrderFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrderFunc1(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOrderFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOrderFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOrderFunc2(tt.args.numCourses, tt.args.prerequisites); !reflect.DeepEqual(got, tt.want) && !reflect.DeepEqual(got, tt.want2) {
				t.Errorf("findOrderFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
