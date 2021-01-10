package main

import (
	"reflect"
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want []string
}{
	{"should be equal", args{[]int{0, 1, 2, 4, 5, 7}}, []string{"0->2", "4->5", "7"}},
	{"should be equal", args{[]int{0, 2, 3, 4, 6, 8, 9}}, []string{"0", "2->4", "6", "8->9"}},
	{"should be equal", args{[]int{}}, []string{}},
	{"should be equal", args{[]int{0}}, []string{"0"}},
}

func Test_summaryRangesFunc1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := summaryRangesFunc1(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("summaryRangesFunc1() = %v, want %v", got, tt.want)
			}
		})
	}
}
