package main

import (
	"reflect"
	"sort"
	"testing"
)

type args struct {
	nums []int
	k    int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be ok", args{[]int{1, 1, 1, 2, 2, 3}, 2}, []int{1, 2}},
	{"should be ok", args{[]int{1}, 1}, []int{1}},
	{"should be ok", args{[]int{1, 2}, 2}, []int{1, 2}},
}

func Test_topKFrequentFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequentFunc1(tt.args.nums, tt.args.k)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequentFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_topKFrequentFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequentFunc2(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequentFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
