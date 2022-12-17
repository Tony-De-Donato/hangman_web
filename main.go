package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", acceuil)
	http.HandleFunc("/jeu", jeu_hangman)
	http.ListenAndServe(":80", nil)

}
func acceuil(w http.ResponseWriter, r *http.Request) {
	tmpl1, err := template.ParseFiles("acceuil_hangman.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl1.Execute(w, nil)
	if r.FormValue("Hangman1") == "1" {
		jeu_hangman(w, r)
	}
}

func jeu_hangman(w http.ResponseWriter, r *http.Request) {
	tmpl2, err := template.ParseFiles("jeu_hangman.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tmpl2.Execute(w, nil)
}
