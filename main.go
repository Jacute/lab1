package main

import (
	"fmt"
	"strconv"
)

func menu() {
	fmt.Println("===MENU===\n1. Rectangle\n2. Triangle\n3. Exit")
}

func menuRectangle() {
	fmt.Println("===RECTANGLE===\n1. Square\n2. Perimeter\n3. Diagonall Length\n4. Back")
}

func menuTriangle() {
	fmt.Println("===Triangle===\n1. Perimeter\n2. Square\n3. Checking for isosceles\n4. Back")
}

func main() {
	for {
		var choice, choice2 int8

		menu()
		fmt.Print("Choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var a, b float64

			fmt.Println("Write sides of rectangle")
			fmt.Print("a: ")
			fmt.Scan(&a)
			fmt.Print("b: ")
			fmt.Scan(&b)

		loop1:
			for {
				menuRectangle()
				fmt.Print("Choice: ")
				fmt.Scan(&choice2)

				switch choice2 {
				case 1:
					fmt.Println("Rectangle Square: " + strconv.FormatFloat(rectangleSquare(a, b), 'f', -1, 64))
				case 2:
					fmt.Println("Rectangle Perimeter: " + strconv.FormatFloat(rectanglePerimeter(a, b), 'f', -1, 64))
				case 3:
					fmt.Println("Diagonall Length: " + strconv.FormatFloat(rectangleDiagonallength(a, b), 'f', -1, 64))
				case 4:
					break loop1
				}
			}

		case 2:
			var a1, b1, c1 float64

			fmt.Println("Write sides of triangle")
			fmt.Print("a: ")
			fmt.Scan(&a1)
			fmt.Print("b: ")
			fmt.Scan(&b1)
			fmt.Print("c: ")
			fmt.Scan(&c1)

		loop2:
			for {
				menuTriangle()
				fmt.Println("Choice: ")
				fmt.Scan(&choice2)

				switch choice2 {
				case 1:
					fmt.Println("Triangle Perimeter: " + strconv.FormatFloat(trianglePerimeter(a1, b1, c1), 'f', -1, 64))
				case 2:
					fmt.Println("Triangle Square: " + strconv.FormatFloat(triangleSquare(a1, b1, c1), 'f', -1, 64))
				case 3:
					if triangleIsIsosceles(a1, b1, c1) {
						fmt.Println("Triangle is isosceles")
					} else {
						fmt.Println("Triangle is not isosceles")
					}
				case 4:
					break loop2
				}
			}

		case 3:
			return
		}
	}
}
