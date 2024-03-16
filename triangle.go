package main

import "math"

func triangleSquare(a float64, b float64, c float64) float64 {
	p := (a + b + c) / 2

	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func trianglePerimeter(a float64, b float64, c float64) float64 {
	return a + b + c
}

func triangleIsIsosceles(a float64, b float64, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}
	return false
}
