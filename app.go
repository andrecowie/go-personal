package main

import (
    	"fmt"
	"log"
//	"io/ioutil"
	"net/http"
	"html/template"
)

type Page struct{
	Title string
	Body []byte
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	p := &Page{Title: "Testing", Body: []byte("This is andre testing")}
	fmt.Printf("%+v\n", r)
	t.Execute(w, p)
}

func main() {
 	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
