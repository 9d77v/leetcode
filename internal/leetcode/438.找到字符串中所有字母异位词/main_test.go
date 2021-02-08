package main

import (
	"reflect"
	"testing"
)

type args struct {
	s string
	p string
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{"cbaebabacd", "abc"}, []int{0, 6}},
	{"should be equal", args{"abab", "ab"}, []int{0, 1, 2}},
}

func Test_findAnagrams(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.args.s, tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
