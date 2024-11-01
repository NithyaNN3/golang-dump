package main

import "fmt"

func main() {
	var name string = "Alice"
	var age int = 30

	var city = "Melbourne" // go infers type from the variable

	number := 42 // shorthand version

	fmt.Println(name, age, city, number)
}
