package day2_test

import (
	"testing"

	"AOC24/internal/day2"
)

func TestRedNosedReports_Part1(t *testing.T) {
	tests := []struct {
		name    string
		rReport day2.RedNosedReports
		want    string
	}{
		{
			name: "Day2 Example",
			rReport: day2.RedNosedReports{
				[]int{7, 6, 4, 2, 1},
				[]int{1, 2, 7, 8, 9},
				[]int{9, 7, 6, 2, 1},
				[]int{1, 3, 2, 4, 5},
				[]int{8, 6, 4, 4, 1},
				[]int{1, 3, 6, 7, 9},
			},
			want: "2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rReport.Part1()
			if got != tt.want {
				t.Errorf("Part1() = %v, want %v", got, tt.want)
			}
		})
	}
}
