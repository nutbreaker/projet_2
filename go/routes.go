package main

import "net/http"

func (a *artBox) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
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
