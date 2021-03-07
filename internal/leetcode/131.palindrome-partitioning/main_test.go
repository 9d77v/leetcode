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
	want [][]string
}{
	{"should be equal", args{"aab"}, [][]string{
		{"a", "a", "b"}, {"aa", "b"},
	}},
}

func Test_partition(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := partition(tt.args.s); !reflect.DeepEqual(gotAns, tt.want) {
				t.Errorf("partition() = %v, want %v", gotAns, tt.want)
			}
		})
	}
}
