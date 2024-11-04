/*
1. Pointers
2. Structs
3. Concurrency
4. Error handling
*/

package main

import "fmt"

func main() {
	// Pointers
	var a *int
	i := 42

	a = &i

	fmt.Println(a)
	fmt.Println(*a)
}
