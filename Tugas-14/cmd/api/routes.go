package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Config) routes() http.Handler {
	router := httprouter.New()

	router.GET("/nilai", app.GetAllNilaiMahasiswa)
	router.POST("/nilai", app.InsertNilaiMahasiswa)

	return router
}