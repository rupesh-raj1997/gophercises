package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type ChapterOptions struct {
	Text string `json:text`
	Arc  string `json:arc`
}

type Chapter struct {
	Title   string           `json:title`
	Story   []string         `json:story`
	Options []ChapterOptions `json:options`
}

type Story map[string]Chapter

func main() {
	// renderContent: function to render content of title, story and options

	// iterate through the story
	// 1. build content for the arc
	// 2. add a route with path for the arc
	// an http handler to take the input of query params

	file, err := os.ReadFile("gopher.json")
	if err != nil {
		panic("Can't load the file")
	}

	var story Story

	err = json.Unmarshal(file, &story)
	if err != nil {
		panic("Can't read the json")
	}

	for k, v := range story {
		// fmt.Println(k, v.Title, v.Story)
		http.HandleFunc("/"+k, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "<h1>%s</h1>", v.Title)
			for _, line := range v.Story {
				fmt.Fprintf(w, "<p>%s</p>", line)
			}
			fmt.Fprintf(w, "<br/><br/><br/>")
			for _, option := range v.Options {
				fmt.Fprintf(w, "<a href='/%s'>%s</a> <br>", option.Arc, option.Text)
			}
		})
	}

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello")
	// })

	log.Fatal(http.ListenAndServe(":8080", nil))
}
