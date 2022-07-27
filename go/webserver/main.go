package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router1 := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, fmt.Sprintf("hello %s", r.URL.Path[:1]))
	}
	home := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	}
	http.HandleFunc("/", home)
	http.HandleFunc("/test", router1)
	port := ":8500"
	fmt.Printf("Running on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
