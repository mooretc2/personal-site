package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Contains data for a blogpost for use in the html template.
type Blogpost struct {
	Title string
	Date string
	Headers []string
	Contents []string
}

var index = template.Must(template.ParseFiles("index.html"))

// loadContent loads site content from text documents.
// It returns a Content struct that contains that information.
func loadContent() ([]Blogpost, error) {

	fileList, err := ioutil.ReadDir("blogposts")
	posts := []Blogpost{}
	if err != nil {
		return nil, err
	}
	for _, file := range fileList {
		//read in and split raw post into sections
		rawPost, err := strings.Split(string(ioutil.ReadFile("blogposts/"+file.Name())), "\n::::\n")
		if err != nil {
			return nil, err
		}
		post := Blogpost{Title: rawPost[0],
			Date: rawPost[1],
			Headers: strings.Split(rawPost[2], "\n::\n")
			Contents: strings.Split(rawPost[3], "\n::\n")}
		posts = append(posts, post)
	}

	return posts
}

// renderTemplate executes the html template using the
// Content struct passed in as input.
func renderTemplate(w http.ResponseWriter, c *Content) {
	err := index.Execute(w, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	content, err := loadContent()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, content)
}

func main() {
	http.HandleFunc("/", handler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":80", nil))
}
