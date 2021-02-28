package main

import (
	"reflect"
	"testing"
)

func Test_getCollisionTimes(t *testing.T) {
	type args struct {
		cars [][]int
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{"should be equal", args{[][]int{
			{1, 2}, {2, 1}, {4, 3}, {7, 2},
		}}, []float64{1.00000, -1.00000, 3.00000, -1.00000}},
		{"should be equal", args{[][]int{
			{3, 4}, {5, 4}, {6, 3}, {9, 1},
		}}, []float64{2.00000, 1.00000, 1.50000, -1.00000}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCollisionTimes(tt.args.cars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCollisionTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
