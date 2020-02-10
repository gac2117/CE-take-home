package main

import "fmt"

func greeting(name string, name2 string) string {
	return "Hello " + name + " and " + name2
}

func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Grace", "Amy"))
	fmt.Println(getSum(3, 4))
}
