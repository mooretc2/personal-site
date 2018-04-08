package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Content is a struct for containing page content
type Content struct {
	Intro      string
	Experience []string
	Projects   []string
	Skills     []string
	Links      []string
	Education  []string
}

var index = template.Must(template.ParseFiles("index.html"))

// loadContent loads site content from text documents.
// It returns a Content struct that contains that information.
func loadContent() (*Content, error) {

	//load content from text files
	i, err := ioutil.ReadFile("content/intro.txt")
	e, err := ioutil.ReadFile("content/experience.txt")
	p, err := ioutil.ReadFile("content/projects.txt")
	s, err := ioutil.ReadFile("content/skills.txt")
	l, err := ioutil.ReadFile("content/links.txt")
	ed, err := ioutil.ReadFile("content/education.txt")
	if err != nil {
		return nil, err
	}

	//convert input to strings, then split content into separate strings
	intro := string(i)
	experience := strings.Split(string(e), "\n::\n")
	projects := strings.Split(string(p), "\n::\n")
	skills := strings.Split(string(s), "\n::\n")
	links := strings.Split(string(l), "\n::\n")
	education := strings.Split(string(ed), "\n::\n")

	return &Content{Intro: intro,
		Experience: experience,
		Projects:   projects,
		Skills:     skills,
		Links:      links,
		Education:  education}, nil
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
