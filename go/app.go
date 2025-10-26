package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"net/url"
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

func (a artBox) isUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func (a artBox) isLengthValid(str string, length int) bool {
	return utf8.RuneCountInString(strings.TrimSpace(str)) > length
}

func (a artBox) templateFuncMaps() template.FuncMap {
	return template.FuncMap{
		"isUrl": a.isUrl,
	}
}
