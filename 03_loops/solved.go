package main

import "fmt"

func main() {
	// Create a for loop that will print out the sum of the numbers between 15 and 25.
	// Hint: declare a variable called 'sum' outside of the for loop to keep track of the sum.

	sum := 0

	for i := 15; i < 26; i++ {
		sum += i
	}

	fmt.Println(sum)

	// Create a for loop that will print out numbers 1 to 30.
	// If the number is divisible by 4, print out the word "FizzBuzz" instead of the number.
	// Hint: Use an if/else statement inside of the for loop.

	for i := 1; i <= 30; i++ {
		if i%4 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}

}
