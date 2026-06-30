package main

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Result string
	Text   string
	Banner string
}

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	hometmpl, _ := template.ParseFiles("templates/home.html")

	hometmpl.Execute(w, nil)
}

func asciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	if banner == "" {
		http.Error(w, "400 Bad Request", http.StatusNotFound)
		return
	}

	result := GenerateArt(text, banner)
	data := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}

	tmpl.Execute(w, data)
}

func switchHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")
	if banner == "" {
		banner = r.FormValue("banner")
	}
	if text == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	result := GenerateArt(text, banner)

	finalResult := PageData{
		Result: result,
		Text:   text,
		Banner: banner,
	}

	tmpl.Execute(w, finalResult)
}
