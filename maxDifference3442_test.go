package main

import "testing"

func TestMaxDifference(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"tzt", -1},
		{"aaaaabbc", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDifference(tt.name); got != tt.want {
				t.Errorf("MaxDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
