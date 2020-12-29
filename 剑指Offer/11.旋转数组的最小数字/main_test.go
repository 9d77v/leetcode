package main

import "testing"

type args struct {
	numbers []int
}

var tests = []struct {
	name    string
	args    args
	wantRes int
}{
	{"should be ok", args{[]int{3, 4, 5, 1, 2}}, 1},
	{"should be ok", args{[]int{2, 2, 2, 0, 1}}, 0},
	{"should be ok", args{[]int{1, 3, 5}}, 1},
	{"should be ok", args{[]int{1}}, 1},
	{"should be ok", args{[]int{3, 1}}, 1},
	{"should be ok", args{[]int{3, 1, 1}}, 1},
	{"should be ok", args{[]int{3, 1, 3}}, 1},
	{"should be ok", args{[]int{1, 3, 3}}, 1},
	{"should be ok", args{[]int{3, 3, 1, 3}}, 1},
	{"should be ok", args{[]int{10, 10, 1, 10, 10}}, 1},
}

func Test_minArray(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := minArray(tt.args.numbers); gotRes != tt.wantRes {
				t.Errorf("minArray() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
