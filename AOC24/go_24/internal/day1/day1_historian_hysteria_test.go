package day1_test

import (
	"testing"

	"AOC24/internal/day1"
)

func TestHistorianList_SolveListDistance(t *testing.T) {
	tests := []struct {
		name  string
		hList day1.HistorianList
		want  int
	}{
		{
			name:  "Day 1 Example",
			hList: day1.HistorianList{X: []int{3, 4, 2, 1, 3, 3}, Y: []int{4, 3, 5, 3, 9, 3}},
			want:  11,
		}, {
			name:  "Zero",
			hList: day1.HistorianList{X: []int{}, Y: []int{}},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hList.SolveListDistance()
			if got != tt.want {
				t.Errorf("SolveListDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHistorianList_SimilarityScore(t *testing.T) {
	tests := []struct {
		name  string
		hList day1.HistorianList
		want  int
	}{
		{
			name:  "Day 1 Example",
			hList: day1.HistorianList{X: []int{3, 4, 2, 1, 3, 3}, Y: []int{4, 3, 5, 3, 9, 3}},
			want:  31,
		}, {
			name:  "Zero",
			hList: day1.HistorianList{X: []int{}, Y: []int{}},
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.hList.SimilarityScore()
			if tt.want != got {
				t.Errorf("SimilarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
