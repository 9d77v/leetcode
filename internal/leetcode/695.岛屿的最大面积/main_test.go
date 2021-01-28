package main

import (
	"testing"
)

type args struct {
	grid [][]int
}

func Test_maxAreaOfIsland(t *testing.T) {
	var tests = []struct {
		name       string
		args       args
		wantResult int
	}{
		{"should be ok", args{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			}}, 6},
		{"should be ok", args{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		}, 0},
		{"should be ok", args{
			[][]int{
				{1},
			},
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaOfIsland(tt.args.grid); got != tt.wantResult {
				t.Errorf("maxAreaOfIsland() = %v, wantResult %v", got, tt.wantResult)
			}
		})
	}
}

func Test_maxAreaOfIslandFunc2(t *testing.T) {
	var tests = []struct {
		name       string
		args       args
		wantResult int
	}{
		{"should be ok", args{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			}}, 6},
		{"should be ok", args{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		}, 0},
		{"should be ok", args{
			[][]int{
				{1},
			},
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := maxAreaOfIslandFunc2(tt.args.grid); gotResult != tt.wantResult {
				t.Errorf("maxAreaOfIslandFunc2() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_maxAreaOfIslandFunc3(t *testing.T) {
	var tests = []struct {
		name       string
		args       args
		wantResult int
	}{
		{"should be ok", args{
			[][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			}}, 6},
		{"should be ok", args{
			[][]int{
				{0, 0, 0, 0, 0, 0, 0, 0},
			},
		}, 0},
		{"should be ok", args{
			[][]int{
				{1},
			},
		}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := maxAreaOfIslandFunc3(tt.args.grid); gotResult != tt.wantResult {
				t.Errorf("maxAreaOfIslandFunc3() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
