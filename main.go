package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.Int("port", 80, "The port to listen on")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Printf("Starting on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
