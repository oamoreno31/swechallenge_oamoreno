package main

import "fmt"

func main() {
	var name = "Omar"

	fmt.Println("Hello, World! ", name, ", nice to meet you.")

	var age int
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)
	//fmt.Println("You are ", age, " years old.")

	if age >= 18 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a minor.")
	}
}
