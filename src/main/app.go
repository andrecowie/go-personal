package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"html/template"
)
type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error {
    filename := "src/text/" + p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := "src/text/" + title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("src/templates/index.html")
	fmt.Printf("%+v", t)
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	p1.save()
	p2, _ := loadPage("TestPage")
	t.Execute(w, p2)
}

func main() {
	fmt.Printf("Starting...")
	http.HandleFunc("/", handler)
	// Render Static Files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/static"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
