package main

import "testing"

type args struct {
	items     [][]string
	ruleKey   string
	ruleValue string
}

var tests = []struct {
	name      string
	args      args
	wantCount int
}{
	{"should be equal", args{[][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "lenovo"},
		{"phone", "gold", "iphone"},
	}, "color", "silver"}, 1},
	{"should be equal", args{[][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "lenovo"},
		{"phone", "gold", "iphone"},
	}, "type", "phone"}, 2},
}

func Test_countMatches(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := countMatches(tt.args.items, tt.args.ruleKey, tt.args.ruleValue); gotCount != tt.wantCount {
				t.Errorf("countMatches() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
