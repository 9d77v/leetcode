package main

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"should be equal", args{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{"should be equal", args{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImage(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flipAndInvertImageFunc2(t *testing.T) {
	type args struct {
		A [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"should be equal", args{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{"should be equal", args{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}}, [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipAndInvertImageFunc2(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImageFunc2() = %v, want %v", got, tt.want)
			}
		})
	}
}
