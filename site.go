package main

import (
	"io/ioutil"
	"html/template"
	"log"
	"net/http"
)

var index = template.Must(template.ParseFiles("index.html"))

func renderTemplate(w http.ResponseWriter){
	err := index.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
