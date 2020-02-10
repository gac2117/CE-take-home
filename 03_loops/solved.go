package main

import "fmt"

func main() {
	// Long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// Short method
	// for i := 1; i <= 10; i++ {
	// 	fmt.Printf("Number %d\n", i)
	// }

	// Fizzbuzz
	// for i := 1; i <= 50; i++ {
	// 	if i%15 == 0 {
	// 		fmt.Println("FizzBuzz")
	// 	} else if i%3 == 0 {
	// 		fmt.Println("Fizz")
	// 	} else if i%5 == 0 {
	// 		fmt.Println("Buzz")
	// 	} else {
	// 		fmt.Println(i)
	// 	}
	// }

	// Odd, Even, Prime
	for i := 1; i <= 50; i++ {
		if i%2 != 0 {
			fmt.Println("Odd number")
		} else if i%10 == 0 {
			fmt.Println("tens")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
