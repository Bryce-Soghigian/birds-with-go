package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
new router func creates the router and
returns it to us
now we can use this function to  create instances of the router and test the router outside of the main function


*/
func newRouter() *mux.Router {
	//Router init.
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")

	//This is the ssr handlers.
	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets", http.FileServer(staticFileDir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	//Handle bird code.
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	return r
}

func main() {
	// The router is now formed by calling the `newRouter` constructor function
	// that we defined above. The rest of the code stays the same
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sup?")
}
