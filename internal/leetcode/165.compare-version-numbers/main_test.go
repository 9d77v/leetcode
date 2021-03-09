package main

import "testing"

func Test_compareVersion(t *testing.T) {
	type args struct {
		version1 string
		version2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"should be equal", args{"1.01", "1.001"}, 0},
		{"should be equal", args{"1.0", "1.0.0"}, 0},
		{"should be equal", args{"1.01", "1.001"}, 0},
		{"should be equal", args{"1.01", "1.001"}, 0},
		{"should be equal", args{"1.01", "1.001"}, 0},
		{"should be equal", args{"1", "1.1"}, -1},
		{"should be equal", args{"1.0", "1"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersion(tt.args.version1, tt.args.version2); got != tt.want {
				t.Errorf("compareVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}
