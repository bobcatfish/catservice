package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/bobcatfish/catservice/cat"
)

func sweetCatBlog(w http.ResponseWriter, r *http.Request) {
	y := cat.Yoshimi()

	tmpl := template.Must(template.ParseFiles("static/index.html"))
	tmpl.Execute(w, y)
}

func main() {
	port := flag.Int("port", 80, "The port to listen on")
	flag.Parse()

	http.HandleFunc("/", sweetCatBlog)
	log.Printf("Starting on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
