package main

import (
	"fmt"
	"http/models/initializertion"
	"net/http"
)

func init() {
	fmt.Println("dddd")
	initializertion.Dotenv()
}

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, Go HTTP!")
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)

}
