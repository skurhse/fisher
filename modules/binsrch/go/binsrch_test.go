package binsrch

import (
	"math/rand"
	"testing"
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

type Case struct {
	name     string
	m        int
	n        int
	points   Points[int]
	abscissa int
	targets  []int
}

func BenchmarkSearch(b *testing.B) {
	rand.Seed(time.Now().UnixNano())

	cases := [2]Case{
		{
			name:    "Small [1k]",
			m:       1e3,
			n:       1e2,
			points:  GeneratePoints[int](1e3),
			targets: GenerateTargets[int](1e3, 1e2),
		},
		{
			name:    "Large [100k]",
			m:       1e5,
			n:       1e2,
			points:  GeneratePoints[int](1e5),
			targets: GenerateTargets[int](1e5, 1e2),
		},
	}

	for i, c := range cases {
		b.Run(c.name, func(b *testing.B) {
			for b.Loop() {
				b.StopTimer()
				t := c.targets[i%c.n]
				b.StartTimer()
				c.points.WhereX(t)
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

func GenerateTargets[C Coordinate](m, n int) []C {
	targets := make([]C, n)

	for i := 0; i < n; i++ {
		targets[i] = C(rand.Intn(m))
	}

	return targets
}
