package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *Config) routes() http.Handler {
	router := httprouter.New()

	// Nilai
	router.GET("/nilai", app.GetAllNilai)
	router.POST("/nilai/create", app.InsertNilai)
	router.PUT("/nilai/:id/update", app.UpdateNilai)
	router.DELETE("/nilai/:id/delete", app.DeleteNilai)

	// Mahasiswa
	router.GET("/mahasiswa", app.GetAllMahasiswa)
	router.POST("/mahasiswa/create", app.InsertMahasiswa)
	router.PUT("/mahasiswa/:id/update", app.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", app.DeleteMahasiswa)

	// Mata Kuliah
	router.GET("/mata-kuliah", app.GetAllMataKuliah)
	router.POST("/mata-kuliah/create", app.InsertMataKuliah)
	router.PUT("/mata-kuliah/:id/update", app.UpdateMataKuliah)
	router.DELETE("/mata-kuliah/:id/delete", app.DeleteMataKuliah)

	return router
}