/*
Concurrency - ability to execute multiple tasks simultaneously
instead of sequentially. Particularly useful for network request tasks or database queries
where you don't need to queue tasks. Also use to parallelize tasks that can be run together.

- goroutines are lightweight threads
*/

package main

import (
	"fmt"
	"time"
)

func printMessage() {
	fmt.Println("Hello from Goroutine")
}

// func main() {
// 	go printMessage()
// 	fmt.Println("Hello from main function")
// 	// Wait for the Goroutine to finish
// 	time.Sleep(time.Second)
// 	// both main()and printMessage() will run at the same time
// }

// Channels in Go give you a way to control and synchronize tasks
func printMessage2(message chan string) {
	time.Sleep(time.Second * 2)
	message <- "Hello from Goroutine"
}

// func main() {
// 	message := make(chan string)
// 	go printMessage2(message)
// 	fmt.Println("Hello from main function")
// 	fmt.Println(<-message)
// }

// Adding 10 goroutines

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
