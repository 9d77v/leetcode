package main

import "testing"

type args struct {
	pushed []int
	popped []int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}}, true},
	{"should be false", args{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}}, false},
}

func Test_validateStackSequences(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateStackSequences(tt.args.pushed, tt.args.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
