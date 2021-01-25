package main

import "testing"

func Test_regionsBySlashes(t *testing.T) {
	type args struct {
		grid []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"should be equal", args{[]string{" /", "/ "}}, 2},
		{"should be equal", args{[]string{" /", "  "}}, 1},
		{"should be equal", args{[]string{"\\/", "/\\"}}, 4},
		{"should be equal", args{[]string{"/\\", "\\/"}}, 5},
		{"should be equal", args{[]string{"//", "/ "}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regionsBySlashes(tt.args.grid); got != tt.want {
				t.Errorf("regionsBySlashes() = %v, want %v", got, tt.want)
			}
		})
	}
}
