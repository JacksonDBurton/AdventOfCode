package day2_test

import (
	"slices"
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

func TestRedNosedReports_Part2(t *testing.T) {
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
			want: "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.rReport.Part2()
			if got != tt.want {
				t.Errorf("Part2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteEle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		list []int
		i    int
		want []int
	}{
		{
			name: "Delete end value",
			list: []int{1, 2, 3},
			i:    2,
			want: []int{1, 2},
		},
		{
			name: "Problem Child",
			list: []int{9, 12, 14, 16, 17, 18, 15},
			i:    6,
			want: []int{9, 12, 14, 16, 17, 18},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day2.DeleteEle(tt.list, tt.i)
			if !slices.Equal(got, tt.want) {
				t.Errorf("DeleteEle() = %v, want %v", got, tt.want)
			}
		})
	}
}
