package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "../database/artbox.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	artBoxApp := newArtBox(db)
	fs := http.FileServer(http.Dir("../img/"))
	artBoxApp.mux.Handle("/img/", http.StripPrefix("/img/", fs))
	artBoxApp.mux.HandleFunc("GET /{$}", artBoxApp.home)
	artBoxApp.mux.HandleFunc("GET /oeuvre/{id}", artBoxApp.oeuvre)
	artBoxApp.mux.HandleFunc("GET /ajouer", artBoxApp.ajouterGet)
	artBoxApp.mux.HandleFunc("POST /ajouer", artBoxApp.ajouterPost)

	server := http.Server{
		Addr:    ":6969",
		Handler: artBoxApp.mux,
	}

	fmt.Printf("Starting server on port %s...\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
