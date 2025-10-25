package main

import (
	"database/sql"
	"net/http"
)

type artBox struct {
	mux *http.ServeMux
	db  *sql.DB
}

func newArtBox(db *sql.DB) *artBox {
	return &artBox{
		mux: http.NewServeMux(),
		db:  db,
	}
}
