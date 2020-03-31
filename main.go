package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sup?")
}
func main() {
	PORT := ":5000"
	http.HandleFunc("/", handler)
	fmt.Println("//=====Listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
