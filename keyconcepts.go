/*
1. Pointers
2. Structs
3. Concurrency
4. Error handling
*/

package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// Pointers
	var a *int
	i := 42

	a = &i

	fmt.Println(a)
	fmt.Println(*a)

	// Structs
	fmt.Println(person{"Bob", 23})
	fmt.Println(person{name: "Alice", age: 23})
	fmt.Println(person{name: "Alice"})
	fmt.Println(&person{name: "Alice", age: 24})
	fmt.Println(newPerson("Jon"))
}
