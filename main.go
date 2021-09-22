package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/bobcatfish/catservice/cat"
)

func sweetCatBlog(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "" || r.URL.Path == "/" {
		tmpl := template.Must(template.ParseFiles("static/index.html"))
		tmpl.Execute(w, struct{Cats []cat.TektonCat}{Cats: cat.GetCatsOfTekton()})
	} else {
		http.ServeFile(w, r, path.Join("static", r.URL.Path))
	}
}

func main() {
	port := flag.Int("port", 80, "The port to listen on")
	flag.Parse()

	http.HandleFunc("/", sweetCatBlog)
	log.Printf("Starting on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
