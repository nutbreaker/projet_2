package main

import (
	"html/template"
	"log"
	"net/http"
)

func (a *artBox) accueil(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./templates/base.html",
		"./templates/accueil.html",
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	oeuvres, err := a.GetOeuvres()
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", map[string]any{
		"oeuvres": oeuvres,
	})
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *artBox) oeuvre(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("oeuvre"))
}

func (a *artBox) ajouterGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ajouter GET"))
}

func (a *artBox) ajouterPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ajouter POST"))
}
