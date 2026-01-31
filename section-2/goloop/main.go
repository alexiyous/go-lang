package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// while loop
	k := 3
	for k > 0 {
		fmt.Println(k)
		k--
	}

	// while infinite (or while true you might say)
	counter := 0
	for {
		fmt.Println("counter", counter)
		counter++
		if counter >= 5 {
			break
		}
	}

	// skipping loop
	for i:= 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	// array loop (tbh i think this is like for each, 
	// and we can access the index & value)
	items := []string{"Go", "Python", "Java"}
	for index, value := range items {
		fmt.Println(index, value)
	}
}