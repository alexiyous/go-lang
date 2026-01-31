package main

import "fmt"

func main() {
	tmp := 25

	if tmp > 30 {
		fmt.Println("greater than 30")
	} else {
		fmt.Println("less than 30")
	}

	score := 85

	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	}

	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	// when you assigns values for map type, it will return you 
	// hasAccess = what is the value of that key 
	// ok = if key "john" exist?
	if hasAccess, ok := userAccess["john"]; ok && hasAccess {
		fmt.Println("Jane has access")
	} else {
		fmt.Println("You don't have access")
	}
}