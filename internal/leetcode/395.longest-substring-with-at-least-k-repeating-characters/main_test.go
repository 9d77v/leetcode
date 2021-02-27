package main

import (
	"testing"
)

func Test_longestSubstring(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"should be equal", args{"aaabb", 3}, 3},
		{"should be equal", args{"ababbc", 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := longestSubstring(tt.args.s, tt.args.k); gotAns != tt.wantAns {
				t.Errorf("longestSubstring() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
