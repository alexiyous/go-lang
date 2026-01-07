package main

import "fmt"

func main() {
	var stringExample string
	stringExample = "Alex"

	fmt.Println(stringExample)

	var count int
	count = 10

	fmt.Println(count)

	var firstName, lastName string
	firstName = "Alex"
	lastName = "Andrianno"

	fmt.Println(firstName, lastName)

	// we can define and initialized with shorter syntax without typing the data type
	age := 22
	fmt.Println(age)

	// we can also type it this way
	var name = "Alex"
	fmt.Println(name)
}
