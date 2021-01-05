package main

import (
	"reflect"
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"should be equal", args{"abbxxxxzzy"}, [][]int{{3, 6}}},
	{"should be equal", args{"abc"}, [][]int{}},
	{"should be equal", args{"abcdddeeeeaabbbcd"}, [][]int{{3, 5}, {6, 9}, {12, 14}}},
	{"should be equal", args{"aba"}, [][]int{}},
	{"should be equal", args{"aaa"}, [][]int{{0, 2}}},
}

func Test_largeGroupPositions(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largeGroupPositions(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largeGroupPositions() = %v, want %v", got, tt.want)
			}
		})
	}
}
