package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Monday"
	fmt.Println("Today is", day)

	switch day {
	case "Sunday", "Saturday":
		fmt.Println("It's Weekend")
	case "Monday":
		fmt.Println("Retro")
	default:
		fmt.Println("Mid-week")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good Morning")
	case hour < 17:
		fmt.Println("Good Afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("Integer: %d\n", t)
		case string:
			fmt.Printf("String: %s\n", t)
		case bool:
			fmt.Printf("Boolean: %t\n", t)
		default:
			fmt.Printf("Unknown type: %T\n", t)
		}
	}

	checkType(21)
	checkType("Test")
	checkType(true)
	checkType(312.23)
}