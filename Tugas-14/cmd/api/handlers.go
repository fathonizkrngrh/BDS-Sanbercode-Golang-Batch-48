package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tugas14/data"
	"tugas14/nilai"
	"tugas14/utils"

	"github.com/julienschmidt/httprouter"
)


func (app *Config) GetAllNilaiMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilaiMhs, err := nilai.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	  utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Success get all nilai mahasiswa",
		Data: nilaiMhs,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) InserNilaiMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") == "application/json" {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var newNilai data.NilaiMahasiswa
		err := json.NewDecoder(r.Body).Decode(&newNilai)
		if err != nil {
			utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
			return
		}

		newNilai.IndeksNilai = GetIndeksNilai(newNilai.Nilai)

		log.Println(newNilai)

		if err := nilai.Insert(ctx, newNilai); err != nil {
			utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
			return
		}

		payload := utils.JsonResponse{
			Code: http.StatusCreated,
			Status: "OK",
			Message: "Success get all nilai mahasiswa",
			Data: newNilai,
		}
	  
		utils.WriteJSON(w, http.StatusCreated, payload)
		
	} else {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		err := r.ParseMultipartForm(0)
		if err != nil {
			utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
			return
		}

		newNilai := data.NilaiMahasiswa{
			Nama:       r.FormValue("nama"),
			MataKuliah: r.FormValue("mata_kuliah"),
			Nilai:      ParseInt(r.FormValue("nilai")),
		}

		newNilai.IndeksNilai = GetIndeksNilai(newNilai.Nilai)

		if err := nilai.Insert(ctx, newNilai); err != nil {
			utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
			return
		}

		payload := utils.JsonResponse{
			Code: http.StatusCreated,
			Status: "OK",
			Message: "Success get all nilai mahasiswa",
			Data: newNilai,
		}
	  
		utils.WriteJSON(w, http.StatusCreated, payload)
	}
}
