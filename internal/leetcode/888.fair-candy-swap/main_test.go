package main

import (
	"reflect"
	"testing"
)

type args struct {
	A []int
	B []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{[]int{1, 1}, []int{2, 2}}, []int{1, 2}},
	{"should be equal", args{[]int{1, 2}, []int{2, 3}}, []int{1, 2}},
	{"should be equal", args{[]int{2}, []int{1, 3}}, []int{2, 3}},
	{"should be equal", args{[]int{1, 2, 5}, []int{2, 4}}, []int{5, 4}},
}

func Test_fairCandySwapFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwapFunc1(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwapFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fairCandySwapFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwapFunc2(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwapFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fairCandySwapFunc3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwapFunc3(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwapFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
