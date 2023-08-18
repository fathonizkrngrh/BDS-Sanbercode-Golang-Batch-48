package data

import (
	"database/sql"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		NilaiMahasiswa: NilaiMahasiswa{},
	}
}

type Models struct {
	NilaiMahasiswa NilaiMahasiswa
}

type (
  NilaiMahasiswa struct {
	ID          int 		`json:"id"`
	Nama        string 		`json:"nama"`
	MataKuliah  string 		`json:"mata_kuliah"`
	IndeksNilai string		`json:"indeks_nilai"`
	Nilai       int   		`json:"nilai"`
    CreatedAt 	time.Time 	`json:"created_at"`
    UpdatedAt 	time.Time 	`json:"updated_at"`
	}
)