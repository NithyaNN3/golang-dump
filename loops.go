package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("the number: ", i)
	}
	// Go doesn't have a while loops explicitly but has a format it follows

	i := 0
	for i < 5 {
		fmt.Println("while numbs: ", i)
		i++
	}

	// Infinite loops

	for {
		fmt.Println("Hellooo")
		// use 'break' to exit or 'return' to end function
		break
	}

	// looping over arrays
	fruits := []string{"apple", "banana"}

	for index, value := range fruits {
		fmt.Println(index, value)
	}

	// looping over maps
	ages := map[string]int{"Alice": 25, "Ella": 27}

	for key, value := range ages {
		fmt.Println(key, value)
	}

	// nested loops
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}
}
