package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "这是一个Go写的Http server")
}

func main() {
	log.Println("Starting http server on localhost:8000")
	http.HandleFunc("/test", handler)
	http.ListenAndServe(":8000", nil)
}
