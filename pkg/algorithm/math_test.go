package algorithm

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
	want []int
}{
	{"should be equal", args{[]int{3, 5, 3, 7, 2, 1}}, []int{1, 2, 3, 3, 5, 7}},
}

func TestQuickSort(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.nums)
			t.Log(tt.args.nums)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Error("sort failed")
			}
		})
	}
}
