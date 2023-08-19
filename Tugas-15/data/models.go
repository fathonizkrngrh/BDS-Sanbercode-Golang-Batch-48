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
		Nilai: Nilai{},
		Mahasiswa: Mahasiswa{},
		MataKuliah: MataKuliah{},
	}
}

type Models struct {
	Nilai Nilai
	Mahasiswa Mahasiswa
	MataKuliah MataKuliah
}

type (
  Nilai struct {
	ID          	int 		`json:"id"`
	Nilai       	int   		`json:"nilai"`
	Indeks 			string		`json:"indeks"`
    CreatedAt 		time.Time 	`json:"created_at"`
    UpdatedAt 		time.Time 	`json:"updated_at"`
	MataKuliahID  	int 		`json:"mata_kuliah_id"`
	MahasiswaID  	int 		`json:"mahasiswa_id"`
	}

  Mahasiswa struct {
	ID          	int 		`json:"id"`
	Nama 			string		`json:"nama"`
    CreatedAt 		time.Time 	`json:"created_at"`
    UpdatedAt 		time.Time 	`json:"updated_at"`
	}

  MataKuliah struct {
	ID          	int 		`json:"id"`
	Nama 			string		`json:"nama"`
    CreatedAt 		time.Time 	`json:"created_at"`
    UpdatedAt 		time.Time 	`json:"updated_at"`
	}


)