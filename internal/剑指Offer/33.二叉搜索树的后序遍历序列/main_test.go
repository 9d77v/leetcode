package main

import (
	"testing"
)

type args struct {
	postorder []int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be false", args{[]int{1, 6, 3, 2, 5}}, false},
	{"should be true", args{[]int{1, 3, 2, 6, 5}}, true},
}

func Test_verifyPostorder(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorder(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyPostorderFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verifyPostorderFunc2(tt.args.postorder); got != tt.want {
				t.Errorf("verifyPostorderFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
