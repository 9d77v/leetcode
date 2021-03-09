package main

import (
	"testing"
)

type args struct {
	S string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"should be equal", args{"abbaca"}, "ca"},
	{"should be equal", args{"azxxzy"}, "ay"},
	{"should be equal", args{"aaaaaaaaa"}, "a"},
}

func Test_removeDuplicates(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.S); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeDuplicatesFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicatesFunc2(tt.args.S); got != tt.want {
				t.Errorf("removeDuplicatesFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
