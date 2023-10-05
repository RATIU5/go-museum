package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.NewServeMux()
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from the Go program"))
	})
	err := http.ListenAndServe(":3333", server)

	if err == nil {
		fmt.Println("Error while opening the server")
	}
}
