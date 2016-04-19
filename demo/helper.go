package demo

import "math"

const (
	tolerance float64 = 1e-10
)

func fEqual(a, b float64) bool {
	return (math.Abs(a-b) < tolerance)
}
