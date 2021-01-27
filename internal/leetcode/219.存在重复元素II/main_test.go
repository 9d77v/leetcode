package main

import "testing"

type args struct {
	nums []int
	k    int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{"should be true", args{[]int{1, 2, 3, 1}, 3}, true},
	{"should be true", args{[]int{1, 0, 1, 1}, 1}, true},
	{"should be true", args{[]int{1, 2, 3, 1, 2, 3}, 2}, false},
}

func Test_containsNearbyDuplicate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
