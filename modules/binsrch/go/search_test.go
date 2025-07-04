package binsrch

import (
	"errors"
	"testing"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		points   []Point
		abscissa int
		lower    int
		upper    int
		err      error
	}{
		// Basic case: abscissa present in the middle
		{
			points:   []Point{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 5,
			lower:    2,
			upper:    3,
			err:      nil,
		},
		// Case: abscissa present at the start
		{
			points:   []Point{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 1,
			lower:    0,
			upper:    1,
			err:      nil,
		},
		// Case: abscissa present at the end
		{
			points:   []Point{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 9,
			lower:    4,
			upper:    5,
			err:      nil,
		},
		// Case: abscissa not present
		{
			points:   []Point{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 6,
			lower:    0,
			upper:    0,
			err:      errors.New("not found"),
		},
		// Case: abscissa present multiple times
		{
			points:   []Point{{1, 2}, {3, 4}, {5, 6}, {5, 7}, {9, 10}},
			abscissa: 5,
			lower:    2,
			upper:    4,
			err:      nil,
		},
		// Case: empty points list
		{
			points:   []Point{},
			abscissa: 1,
			lower:    0,
			upper:    0,
			err:      errors.New("points empty"),
		},
	}

	for _, tt := range tests {
		lower, upper, err := Search(tt.points, tt.abscissa)
		if lower != tt.lower || upper != tt.upper || (err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("Search(%v, %d) = (%d, %d, %v), want (%d, %d, %v)",
				tt.points, tt.abscissa, lower, upper, err, tt.lower, tt.upper, tt.err)
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	testCases := []struct {
		name     string
		points   []Point
		abscissa int
	}{
		{
			name:     "Small",
			points:   generatePoints(1e6),
			abscissa: 5e5,
		},
		{
			name:     "Large",
			points:   generatePoints(1e9),
			abscissa: 5e8,
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, err := Search(tc.points, tc.abscissa)
				if err != nil && err.Error() != "not found" {
					b.Fatalf("unexpected error: %v", err)
				}
			}
		})
	}
}

func generatePoints(n int) []Point {
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = Point{X: i, Y: i}
	}
	return points
}
