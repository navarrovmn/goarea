package goarea

import "math"

// Numeric proportion between circle and diameter
const Pi = 3.1416

// Circ area
func Circ(r float64) float64 {
	return math.Pow(r, 2) * Pi
}

//  Rect area
func Rect(b, h float64) float64 {
	return b * h;
}
