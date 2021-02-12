package main

import (
	"reflect"
	"testing"
)

type args struct {
	rowIndex int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{"should be equal", args{0}, []int{1}},
	{"should be equal", args{1}, []int{1, 1}},
	{"should be equal", args{2}, []int{1, 2, 1}},
	{"should be equal", args{3}, []int{1, 3, 3, 1}},
	{"should be equal", args{4}, []int{1, 4, 6, 4, 1}},
}

func Test_getRow(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRowFunc2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRowFunc2(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRowFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
