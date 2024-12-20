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

// func main() {
// 	for i := 0; i < 10; i++ {
// 		go f(i)
// 	}
// 	var input string
// 	fmt.Scanln(&input)
// }

// Channels - provide a way in which goroutines can communicate with each other
func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" // c <- "ping" means send "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c // msg := <- c means receive a message and store it in msg
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// func main() {
// 	var c chan string = make(chan string)

// 	go pinger(c) // When pinger attempts to send a message on the channel
// 	// it will wait until printer is ready to receive the message. (this is known as blocking)
// 	go ponger(c)
// 	go printer(c)

// 	var input string
// 	fmt.Scanln(&input)
// }

/*
Channel direction can be written in this format:
func pinger(c chan<- string)
or
func printer(c <-chan string)
*/

/* Buffered channels -
are channels that can send multiple values without requiring a corresponding receive for each value. This is unlike regular channels where
each send needs a receive
*/

// Example - make(chan int, 5)

func main() {
	// Create a buffered channel with capacity of 3
	ch := make(chan int, 3)

	// Send values to the channel
	ch <- 1
	ch <- 2
	ch <- 3

	// The channel is now full, so any additional sends would block
	// Receiving values
	fmt.Println(<-ch) // Output: 1
	fmt.Println(<-ch) // Output: 2
	fmt.Println(<-ch) // Output: 3
}
