package main

import "math"

func rectanglePerimeter(a float64, b float64) float64 {
	return (a + b) * 2
}

func rectangleSquare(a float64, b float64) float64 {
	return a * b
}

func rectangleDiagonallength(a float64, b float64) float64 {
	return math.Sqrt(a*a + b*b)
}
