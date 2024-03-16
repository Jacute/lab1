package main

import (
	"fmt"
	"strconv"
)

func menu() {
	fmt.Println(green("===") + cyan("MENU") + green("===\n") + "1. Rectangle\n2. Triangle\n3. Exit")
}

func menuRectangle() {
	fmt.Println(green("===") + cyan("RECTANGLE") + green("===\n") + "1. Square\n2. Perimeter\n3. Diagonall Length\n4. Back")
}

func menuTriangle() {
	fmt.Println(green("===") + cyan("Triangle") + green("===\n") + "1. Perimeter\n2. Square\n3. Checking for isosceles\n4. Back")
}

func main() {
	for {
		var choice, choice2 int8

		menu()
		fmt.Print(cyan("Choice: "))
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
				fmt.Print(cyan("Choice: "))
				fmt.Scan(&choice2)

				switch choice2 {
				case 1:
					fmt.Println("Rectangle Square: " + green(strconv.FormatFloat(rectangleSquare(a, b), 'f', -1, 64)))
				case 2:
					fmt.Println("Rectangle Perimeter: " + green(strconv.FormatFloat(rectanglePerimeter(a, b), 'f', -1, 64)))
				case 3:
					fmt.Println("Diagonall Length: " + green(strconv.FormatFloat(rectangleDiagonallength(a, b), 'f', -1, 64)))
				case 4:
					break loop1
				}
			}

		case 2:
			var a, b, c float64

			fmt.Println("Write sides of triangle")
			fmt.Print("a: ")
			fmt.Scan(&a)
			fmt.Print("b: ")
			fmt.Scan(&b)
			fmt.Print("c: ")
			fmt.Scan(&c)

			if !isTriangle(a, b, c) {
				fmt.Println(red("There is no triangle with this sides"))
				break
			}

		loop2:
			for {
				menuTriangle()
				fmt.Println("Choice: ")
				fmt.Scan(&choice2)

				switch choice2 {
				case 1:
					fmt.Println("Triangle Perimeter: " + green(strconv.FormatFloat(trianglePerimeter(a, b, c), 'f', -1, 64)))
				case 2:
					fmt.Println("Triangle Square: " + green(strconv.FormatFloat(triangleSquare(a, b, c), 'f', -1, 64)))
				case 3:
					if triangleIsIsosceles(a, b, c) {
						fmt.Println(green("Triangle is isosceles"))
					} else {
						fmt.Println(red("Triangle is not isosceles"))
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
