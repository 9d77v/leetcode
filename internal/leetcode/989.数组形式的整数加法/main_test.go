package main

import (
	"reflect"
	"testing"
)

type args struct {
	A []int
	K int
}

var tests = []struct {
	name       string
	args       args
	wantResult []int
}{
	{"should be equal", args{[]int{1, 2, 0, 0}, 34}, []int{1, 2, 3, 4}},
	{"should be equal", args{[]int{2, 7, 4}, 181}, []int{4, 5, 5}},
	{"should be equal", args{[]int{2, 1, 5}, 806}, []int{1, 0, 2, 1}},
	{"should be equal", args{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1}, []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},
}

func Test_addToArrayForm(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := addToArrayForm(tt.args.A, tt.args.K); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("addToArrayForm() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
