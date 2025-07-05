package main

import (
	"fmt"
	"net/http"
)

// handler funciton
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellwo world")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm hasib, a software engineer learning golang")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "email : mahabubulhasib@gmail.com\nphone : 01956832167\nif you relly want to contact me try several things at all otherwise I don't see any hope actually about your success")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/contact", contactHandler)

	fmt.Println("Server runngin on : 3000")

	err := http.ListenAndServe(":3000", mux)

	if err != nil {
		fmt.Println("failed to start the server", err)
	}
}
