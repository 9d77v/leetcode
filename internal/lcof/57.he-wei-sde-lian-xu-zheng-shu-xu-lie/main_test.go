package main

import (
	"reflect"
	"testing"
)

type args struct {
	target int
}

var tests = []struct {
	name string
	args args
	want [][]int
}{
	{"should be equal", args{9}, [][]int{
		{2, 3, 4}, {4, 5},
	}},
	{"should be equal", args{15}, [][]int{
		{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8},
	}},
	{"should be equal", args{10}, [][]int{
		{1, 2, 3, 4},
	}},
}

func Test_findContinuousSequence(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuousSequence(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findContinuousSequenceFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findContinuousSequenceFunc2(tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findContinuousSequenceFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
