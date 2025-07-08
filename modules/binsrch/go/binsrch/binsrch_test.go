package binsrch

import (
	"testing"
	"math/rand"
	"time"
)

func TestSearch(t *testing.T) {
	tests := []struct {
		points   Points[int]
		abscissa int
		lower    int
		upper    int
		ok       bool
	}{
		// Basic case: abscissa present in the middle
		{
			points:   Points[int]{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 5,
			lower:    2,
			upper:    2,
			ok:       true,
		},
		// Case: abscissa present at the start
		{
			points:   Points[int]{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 1,
			lower:    0,
			upper:    0,
			ok:       true,
		},
		// Case: abscissa present at the end
		{
			points:   Points[int]{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 9,
			lower:    4,
			upper:    4,
			ok:       true,
		},
		// Case: abscissa not present
		{
			points:   Points[int]{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}},
			abscissa: 0,
			ok:       false,
		},
		// Case: abscissa present multiple times
		{
			points:   Points[int]{{1, 2}, {3, 4}, {5, 6}, {5, 7}, {9, 10}},
			abscissa: 5,
			lower:    2,
			upper:    3,
			ok:       true,
		},
		// Case: empty points list
		{
			points:   Points[int]{},
			abscissa: 1,
			ok:       false,
		},
	}

	for _, test := range tests {
		actual, ok := test.points.WhereX(test.abscissa)

		if ok != test.ok {
			t.Errorf("got %v, %v -> %v, want %v", test.points, test.abscissa, ok, test.ok)
		}

		if ok == false {
			if actual != nil {
				t.Errorf("got %v, want nil", actual)
			}
			continue
		}

		expected := test.points[test.lower : test.upper+1]

		for index, expectedPoint := range expected {
			if len(actual) <= index {
				t.Errorf("got %v -> %v, want %v -> %v", test.points, actual, test.points, expected)
				break
			}
			actualPoint := actual[index]
			if expectedPoint != actualPoint {
				t.Errorf("got %v -> %v, want %v -> %v", test.points, actual, test.points, expected)
				break
			}
		}
	}
}

func BenchmarkSearch(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	testCases := []struct {
		name     string
		points   Points[int]
		abscissa int
	}{
		{
			name:     "Small",
			points:   GeneratePoints[int](1e6),
			abscissa: rand.Intn(1e6),
		},
		{
			name:     "Large",
			points:   GeneratePoints[int](1e9),
			abscissa: rand.Intn(1e9),
		},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tc.points.WhereX(tc.abscissa)
			}
		})
	}
}

func GeneratePoints[C Coordinate](n int) Points[C] {
	points := make(Points[C], n)
	for i := 0; i < n; i++ {
		points[i] = Point[C]{C(i), C(i)}
	}
	return points
}
