package day1_test

import (
	"testing"

	"AOC24/internal/day1"
)

func TestSolveListDistance(t *testing.T) {
	tests := []struct {
		name string
		x    []int
		y    []int
		want int
	}{
		{
			name: "Day 1 Example",
			x:    []int{3, 4, 2, 1, 3, 3},
			y:    []int{4, 3, 5, 3, 9, 3},
			want: 11,
		},
		{
			name: "Zero",
			x:    []int{},
			y:    []int{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day1.SolveListDistance(tt.x, tt.y)
			if got != tt.want {
				t.Errorf("SolveListDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
