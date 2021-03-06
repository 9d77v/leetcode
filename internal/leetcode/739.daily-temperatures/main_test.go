package main

import (
	"reflect"
	"testing"
)

type args struct {
	T []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"test ok", args{[]int{73, 74, 75, 71, 69, 72, 76, 73}}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
}

func Test_dailyTemperatures(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dailyTemperatures(tt.args.T); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
