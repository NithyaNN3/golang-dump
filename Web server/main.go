package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n") // Message logged on the server console
	io.WriteString(w, "This is my website!\n")  // response
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")  // Message logged on the server console
	io.WriteString(w, "Hello, HTTP\n")  // response 
}

func main() {
	http.HandleFunc("/", getRoot) // handler registration, maps / to getRoot and /hello to getHello
	http.HandleFunc("/hello", getHello)

	err := http.ListenAndServe(":3333", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
