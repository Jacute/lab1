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

func main() {
	for {
		var choice, choice2 int8

		menu()
		fmt.Print("Choice: ")
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			var a, b float64

			fmt.Println("Write sides of rectangle")
			fmt.Print("a: ")
			fmt.Scanf("%f", &a)
			fmt.Print("b: ")
			fmt.Scanf("%f", &b)

		loop:
			for {
				menuRectangle()
				fmt.Print("Choice: ")
				fmt.Scanf("%d", &choice2)

				switch choice2 {
				case 1:
					fmt.Println("Rectangle Square: " + strconv.FormatFloat(rectangleSquare(a, b), 'f', -1, 64))
				case 2:
					fmt.Println("Rectangle Perimeter: " + strconv.FormatFloat(rectanglePerimeter(a, b), 'f', -1, 64))
				case 3:
					fmt.Println("Diagonall Length: " + strconv.FormatFloat(rectangleDiagonallength(a, b), 'f', -1, 64))
				case 4:
					break loop
				}
			}
		case 2:
		case 3:
			return
		}
	}
}
