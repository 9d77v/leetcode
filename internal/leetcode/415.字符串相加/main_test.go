package main

import "testing"

type args struct {
	num1 string
	num2 string
}

var tests = []struct {
	name string
	args args
	want string
}{
	{"should be equal", args{"0", "0"}, "0"},
}

func Test_addStrings(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addStrings(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("addStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
