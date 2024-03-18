package main

import "math"

func isTriangle(a, b, c float64) bool {
	return a+b > c && a+c > b && b+c > a
}

func triangleSquare(a, b, c float64) float64 {
	p := (a + b + c) / 2

	return math.Sqrt(p * (p - a) * (p - b) * (p - c))
}

func trianglePerimeter(a, b, c float64) float64 {
	return a + b + c
}

func triangleIsIsosceles(a, b, c float64) bool {
	if a == b || a == c || b == c {
		return true
	}
	return false
}
