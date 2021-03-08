package algorithm

import "testing"

func TestIsPalidrome(t *testing.T) {
	type args struct {
		s     string
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"should be true", args{"abb", 1, 2}, true},
		{"should be false", args{"abb", 0, 2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalidrome(tt.args.s, tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("IsPalidrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
