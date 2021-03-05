package main

import (
	"reflect"
	"testing"
)

func Test_singleNumbers(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{[]int{4, 1, 4, 6}}, []int{6, 1}},
		{"should be equal", args{[]int{1, 2, 10, 4, 1, 4, 3, 3}}, []int{2, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumbers(tt.args.nums)
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
