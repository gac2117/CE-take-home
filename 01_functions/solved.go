package main

import (
	"fmt"
	"strconv"
)

// Create a function called 'greeting' that will take in three parameters: name, city, and job title.
// The function must return a string that says "Hello, my name is ______. I live in ______ and I am a ______."
func greeting(name, city, job string) string {
	return "Hello, my name is " + name + ". I live in " + city + " and I am a " + job + "."
}

// Create a function called 'age' that will take in one parameter: birthyear.
// The function must return a string that says "I am _____ years old."
// In order to convert the integer into a string, you must use strconv.Itoa(num). For example, i := strconv.Itoa(10) will return a string of "10".

func age(birthyear int) string {
	age := strconv.Itoa(2020 - birthyear)
	return "I am " + age + " years old."
}

// Call the function 'greeting' and 'age' three times using different parameters.
func main() {
	fmt.Println(greeting("John", "New York City", "banker"))
	fmt.Println(age(1980))
	fmt.Println(greeting("Sam", "Miami", "surfer"))
	fmt.Println(age(1991))
	fmt.Println(greeting("Jane", "San Francisco", "software developer"))
	fmt.Println(age(1986))
}
