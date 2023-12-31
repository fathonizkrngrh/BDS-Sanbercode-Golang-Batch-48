package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type NilaiMahasiswa struct {
	ID          int `json:"id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string
	Nilai       int   `json:"nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}
var idCounter int = 1

func HandlerNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodPost || r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			var newNilai NilaiMahasiswa
			err := json.NewDecoder(r.Body).Decode(&newNilai)
			if err != nil {
				WriteResponse(w, http.StatusBadRequest, "Bad Request", err.Error(), nil)
				return
			}

			newNilai.ID = idCounter
			idCounter++
			newNilai.IndeksNilai = GetIndeksNilai(newNilai.Nilai)
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)
			WriteResponse(w, http.StatusOK, "SUCCESS", fmt.Sprintf("Success create nilai mahasiswa %s", newNilai.Nama), newNilai)
		} else {
			err := r.ParseMultipartForm(0)
			if err != nil {
				WriteResponse(w, http.StatusBadRequest, "Bad Request", err.Error(), nil)
				return
			}

			newNilai := NilaiMahasiswa{
				Nama:       r.FormValue("nama"),
				MataKuliah: r.FormValue("mata_kuliah"),
				Nilai:      ParseInt(r.FormValue("nilai")),
			}

			newNilai.ID = idCounter
			idCounter++
			newNilai.IndeksNilai = GetIndeksNilai(newNilai.Nilai)
			nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)
			WriteResponse(w, http.StatusOK, "SUCCESS", fmt.Sprintf("Success create nilai mahasiswa %s", newNilai.Nama), newNilai)
		}
	} else if r.Method == http.MethodGet {
		totalMhs := len(nilaiNilaiMahasiswa)
		WriteResponse(w, http.StatusOK, "SUCCESS", fmt.Sprintf("Success get all %d nilai mahasiswa ", totalMhs), nilaiNilaiMahasiswa)
	} else  {
		WriteResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", errors.New("Your method doesnt allowed").Error(), nil)
	}
}

