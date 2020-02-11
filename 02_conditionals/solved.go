package main

import "fmt"

func main() {
	z := 7

	if z == 5 {
		fmt.Println("z is equal to 5")
	} else {
		fmt.Println("z is not equal to 5")
	}

	x := 10
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else if x < 10 {
		fmt.Println("x is less than 10")
	} else {
		fmt.Println("x is equal to 10")
	}

	color := "green"

	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is not blue or red")
	}

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}
}
