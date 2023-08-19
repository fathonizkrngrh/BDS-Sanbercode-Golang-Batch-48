package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"tugas15/form"
	"tugas15/mahasiswa"
	"tugas15/mata_kuliah"
	"tugas15/nilai"
	"tugas15/utils"

	"github.com/julienschmidt/httprouter"
)

// Nilai
func (app *Config) GetAllNilai(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	nilaiMhs, err := nilai.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	  utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
	  return
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Success get all nilai mahasiswa",
		Data: nilaiMhs,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) InsertNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var newNilai form.InsertNilai
	err := json.NewDecoder(r.Body).Decode(&newNilai)
	if err != nil {
		utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
	}

	if newNilai.Nilai > 100 {
		utils.ErrorJSON(w, errors.New("Nilai tidak boleh melebihi 100 !!"), "BAD REQUEST",http.StatusBadRequest)
		return
	}

	isMahasiswaExist, err := mahasiswa.IsExist(ctx, newNilai.MahasiswaID)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if !isMahasiswaExist {
		utils.ErrorJSON(w, errors.New("Mahasiswa ID tidak tersedia"), "BAD REQUEST", http.StatusBadRequest)
		return
	}

	isMataKuliahExist, err := mata_kuliah.IsExist(ctx, newNilai.MataKuliahID)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if !isMataKuliahExist {
		utils.ErrorJSON(w, errors.New("Mata Kuliah ID tidak tersedia"), "BAD REQUEST", http.StatusBadRequest)
		return
	}

	isDuplicate, err := nilai.IsDuplicateNilai(ctx, newNilai.MahasiswaID, newNilai.MataKuliahID)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if isDuplicate {
		utils.ErrorJSON(w, errors.New("Data dengan Nama dan Mata Kuliah tersebut sudah ada"), "CONFLICT", http.StatusConflict)
		return
	}

	newNilai.Indeks = GetIndeksNilai(newNilai.Nilai)

	if err := nilai.Insert(ctx, newNilai); err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success insert nilai mahasiswa %v", newNilai.MahasiswaID),
		Data: newNilai,
	}
	
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (app *Config) UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var newNilai form.UpdateNilai

    if err := json.NewDecoder(r.Body).Decode(&newNilai); err != nil {
        utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
    }

	if newNilai.Nilai > 100 {
		utils.ErrorJSON(w, errors.New("Nilai tidak boleh melebihi 100 !!"), "BAD REQUEST",http.StatusBadRequest)
		return
	}

    var idNilai = ps.ByName("id")

	newNilai.Indeks = GetIndeksNilai(newNilai.Nilai)

    if err := nilai.UpdateByID(ctx, newNilai, ParseInt(idNilai)); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success update nilai mahasiswa %v", newNilai.MahasiswaID),
		Data: newNilai,
	}
  
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (app *Config) DeleteNilai(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idNilai = ParseInt(ps.ByName("id"))

    if err := nilai.DeleteByID(ctx, idNilai); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success delete nilai mahasiswa %v", idNilai),
	}
  
	utils.WriteJSON(w, http.StatusCreated, payload)
}

// Mahasiswa
func (app *Config) GetAllMahasiswa(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mhs, err := mahasiswa.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	  utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Success get all mahasiswa",
		Data: mhs,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) InsertMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var newMhs form.InsertMahasiswa
	err := json.NewDecoder(r.Body).Decode(&newMhs)
	if err != nil {
		utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
	}

	isDuplicate, err := mahasiswa.IsDuplicateMahasiswa(ctx, newMhs.Nama)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if isDuplicate {
		utils.ErrorJSON(w, errors.New("Data mahasiswa dengan Nama tersebut sudah ada"), "CONFLICT", http.StatusConflict)
		return
	}

	if err := mahasiswa.Insert(ctx, newMhs); err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success insert mahasiswa %v", newMhs.Nama),
		Data: newMhs,
	}
	
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (app *Config) UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var newMhs form.UpdateMahasiswa

    if err := json.NewDecoder(r.Body).Decode(&newMhs); err != nil {
        utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
    }

    var idMhs = ps.ByName("id")

    if err := mahasiswa.UpdateByID(ctx, newMhs, ParseInt(idMhs)); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success update mahasiswa %s", newMhs.Nama),
		Data: newMhs,
	}
  
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (app *Config) DeleteMahasiswa(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idMhs = ParseInt(ps.ByName("id"))

    if err := mahasiswa.DeleteByID(ctx, idMhs); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success delete mahasiswa %v", idMhs),
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

// Mata Kuliah
func (app *Config) GetAllMataKuliah(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	matkul, err := mata_kuliah.GetAll(ctx)
  
	if err != nil {
	  fmt.Println(err)
	  utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
	}

	payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: "Success get all mata kuliah",
		Data: matkul,
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) InsertMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var newMatkul form.InsertMataKuliah
	err := json.NewDecoder(r.Body).Decode(&newMatkul)
	if err != nil {
		utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
	}

	isDuplicate, err := mata_kuliah.IsDuplicateMataKuliah(ctx, newMatkul.Nama)
	if err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}

	if isDuplicate {
		utils.ErrorJSON(w, errors.New("Data mahasiswa dengan Nama tersebut sudah ada"), "CONFLICT", http.StatusConflict)
		return
	}

	if err := mata_kuliah.Insert(ctx, newMatkul); err != nil {
		utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
	}

	payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success insert mata kuliah %v", newMatkul.Nama),
		Data: newMatkul,
	}
	
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (app *Config) UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    if r.Header.Get("Content-Type") != "application/json" {
        http.Error(w, "Gunakan content type application / json", http.StatusBadRequest)
        return
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var newMatkul form.UpdateMataKuliah

    if err := json.NewDecoder(r.Body).Decode(&newMatkul); err != nil {
        utils.ErrorJSON(w, err, "BAD REQUEST",http.StatusBadRequest)
		return
    }

    var idMatkul = ps.ByName("id")

    if err := mata_kuliah.UpdateByID(ctx, newMatkul, ParseInt(idMatkul)); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusCreated,
		Status: "CREATED",
		Message: fmt.Sprintf("Success update mata kuliah %s", newMatkul.Nama),
		Data: newMatkul,
	}
  
	utils.WriteJSON(w, http.StatusCreated, payload)
}

func (app *Config) DeleteMataKuliah(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    var idMatkul = ParseInt(ps.ByName("id"))

    if err := mata_kuliah.DeleteByID(ctx, idMatkul); err != nil {
        utils.ErrorJSON(w, err, "INTERNAL SERVER ERROR",http.StatusInternalServerError)
		return
    }

    payload := utils.JsonResponse{
		Code: http.StatusOK,
		Status: "OK",
		Message: fmt.Sprintf("Success delete mata kuliah %v", idMatkul),
	}
  
	utils.WriteJSON(w, http.StatusOK, payload)
}