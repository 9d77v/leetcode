package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_getLeastNumbers(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{[]int{0, 0, 0, 2, 0, 5}, 0}, []int{}},
		{"should be equal", args{[]int{3, 2, 1}, 2}, []int{1, 2}},
		{"should be equal", args{[]int{0, 1, 2, 1}, 1}, []int{0}},
		{"should be equal", args{[]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8}, []int{0, 0, 1, 1, 2, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLeastNumbers(tt.args.arr, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLeastNumbersFunc2(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{[]int{0, 0, 0, 2, 0, 5}, 0}, []int{}},
		{"should be equal", args{[]int{3, 2, 1}, 2}, []int{1, 2}},
		{"should be equal", args{[]int{0, 1, 2, 1}, 1}, []int{0}},
		{"should be equal", args{[]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8}, []int{0, 0, 1, 1, 2, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLeastNumbersFunc2(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbersFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLeastNumbersFunc3(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"should be equal", args{[]int{0, 0, 0, 2, 0, 5}, 0}, []int{}},
		{"should be equal", args{[]int{3, 2, 1}, 2}, []int{1, 2}},
		{"should be equal", args{[]int{0, 1, 2, 1}, 1}, []int{0}},
		{"should be equal", args{[]int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}, 8}, []int{0, 0, 1, 1, 2, 2, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getLeastNumbersFunc3(tt.args.arr, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLeastNumbersFunc3() = %v, want %v", got, tt.want)
			}
		})
	}
}
