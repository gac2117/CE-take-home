package main

import "fmt"

func main() {
	// Create a if/else statement that will take in one parameter of a number. Declare and assign the variable.
	// If the number is an even number, print out the message "This is an even number".
	// If the number is an odd number, print out the message "This is an odd number".

	x := 51

	if x%2 == 0 {
		fmt.Println("This is an even number")
	} else {
		fmt.Println("This is an odd number")
	}

	// Create a else-if statement that will take in one parameter of a number between 0 and 23, representing a 24-hour clock. Declare and assign the variable.
	// If the number is between 6 and 12, print out the message "Good morning!"
	// If the number is between 13 and 18, print out the message "Good afternoon!"
	// If the number is either between 0 and 5 or 19 and 23, print out the message "Good night!"
	// If the number is above 23, print out the message "Please use a number between 0 and 23"

	hour := 25

	if hour >= 6 && hour <= 12 {
		fmt.Println("Good morning!")
	} else if hour >= 13 && hour <= 18 {
		fmt.Println("Good afternoon!")
	} else if hour >= 24 {
		fmt.Println("Please use a number between 0 and 23")
	} else {
		fmt.Println("Good night!")
	}

	// Create a switch statement that will take in one parameter of a string, that is a language name. Decare and assign the variable.
	// Choose three languages that you know how to say "Hello" in. You can use Google Translator if needed.
	// Create a case for each of the three language and have it print out "Hello!" in that language.
	// The default case can return "Hello!" in English.

	language := "French"

	switch language {
	case "French":
		fmt.Println("Bonjour!")
	case "Korean":
		fmt.Println("An-nyoung!")
	case "Japanese":
		fmt.Println("Konnichiwa!")
	default:
		fmt.Println("Hello!")
	}
}
