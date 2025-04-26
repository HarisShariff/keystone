package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting Keystone on port 8080...")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Keystone is healthy.")
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
