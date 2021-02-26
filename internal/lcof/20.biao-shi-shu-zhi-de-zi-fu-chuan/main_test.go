package main

import (
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{" "}, false},
	{"should be true", args{" 1 "}, true},
	{"should be true", args{"+100"}, true},
	{"should be true", args{"5e2"}, true},
	{"should be true", args{"-123"}, true},
	{"should be true", args{"3.1416"}, true},
	{"should be true", args{"-1E-16"}, true},
	{"should be true", args{"0123"}, true},
	{"should be false", args{"3."}, true},
	{"should be false", args{"12e"}, false},
	{"should be false", args{"1a3.14"}, false},
	{"should be false", args{"1.2.3"}, false},
	{"should be false", args{"+-5"}, false},
	{"should be false", args{"12e+5.4"}, false},
}

func Test_isNumberFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumberFunc1(tt.args.s); got != tt.want {
				t.Errorf("isNumberFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumberFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isNumberFunc2(tt.args.s); got != tt.want {
				t.Errorf("isNumberFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
