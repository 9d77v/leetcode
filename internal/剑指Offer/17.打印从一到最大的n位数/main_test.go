package main

import (
	"reflect"
	"testing"
)

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"shoule be ok", args{1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
}

func Test_printNumbersFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbersFunc1(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printNumbersFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printNumbersFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printNumbersFunc2(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printNumbersFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
