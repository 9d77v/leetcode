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
	want []string
}{
	{"should be ok", args{"abc"}, []string{"abc", "acb", "bac", "bca", "cba", "cab"}},
}

func Test_permutation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
