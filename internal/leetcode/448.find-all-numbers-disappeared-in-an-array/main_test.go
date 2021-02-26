package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

func Test_findDisappearedNumbers(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbers(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findDisappearedNumbersFunc2(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{[]int{4, 3, 2, 7, 8, 2, 3, 1}}, []int{5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDisappearedNumbersFunc2(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDisappearedNumbersFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
