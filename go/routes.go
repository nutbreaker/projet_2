package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (a *artBox) accueil(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("accueil").Funcs(a.templateFuncMaps()).ParseFiles(
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
	tmpl, err := template.New("oeuvre").Funcs(a.templateFuncMaps()).ParseFiles(
		"./templates/base.html",
		"./templates/oeuvre.html",
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	id, err := strconv.ParseInt(r.PathValue("id"), 10, 0)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	oeuvre, err := a.GetOeuvreById(int(id))
	if err != nil {
		log.Print(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", oeuvre)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *artBox) ajouterGet(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./templates/base.html",
		"./templates/ajouter.html",
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (a *artBox) ajouterPost(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(
		"./templates/base.html",
		"./templates/ajouter.html",
	)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	o := oeuvre{
		Titre:       r.PostFormValue("titre"),
		Description: r.PostFormValue("description"),
		Artiste:     r.PostFormValue("artiste"),
		Image:       r.PostFormValue("image"),
	}

	data := map[string]any{
		"oeuvre": o,
		"error": map[string]bool{
			"titre":       !a.isLengthValid(o.Titre, 0),
			"description": !a.isLengthValid(o.Description, 2),
			"artiste":     !a.isLengthValid(o.Artiste, 0),
			"image":       !a.isUrl(o.Image),
		},
	}

	isFormValid := true
	for _, value := range data["error"].(map[string]bool) {
		isFormValid = isFormValid && !value
	}

	if isFormValid && a.addOeuvre(o) == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	if err = tmpl.ExecuteTemplate(w, "base", data); err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
