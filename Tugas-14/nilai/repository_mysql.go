package nilai

import (
	"context"
	"fmt"
	"log"
	"time"
	"tugas14/config"
	"tugas14/data"
)

const (
	table          = "nilai_mhs"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]data.NilaiMahasiswa, error) {
	var nilaiMhs []data.NilaiMahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai data.NilaiMahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Nama,
			&nilai.MataKuliah,
			&nilai.Nilai,
			&nilai.IndeksNilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		nilai.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		nilai.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		nilaiMhs = append(nilaiMhs, nilai)
	}
	return nilaiMhs, nil
}

func IsDuplicateMahasiswa(ctx context.Context, nama, mataKuliah string) (bool, error) {
    db, err := config.MySQL()
    if err != nil {
        return false, err
    }

    queryText := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE nama = '%v' AND mata_kuliah = '%v'", table, nama, mataKuliah)
    var count int
    err = db.QueryRowContext(ctx, queryText).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func Insert(ctx context.Context, nilai data.NilaiMahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
	  log.Fatal("Can't connect to MySQL", err)
	}
	
	queryText := fmt.Sprintf("INSERT IGNORE INTO %v (nama, mata_kuliah, indeks_nilai, nilai, created_at, updated_at) values('%v','%v','%v',%v, NOW(), NOW())", table,
	  	nilai.Nama,
	  	nilai.MataKuliah,
		nilai.IndeksNilai,
		nilai.Nilai,)
	_, err = db.ExecContext(ctx, queryText)
  
	if err != nil {
	  return err
	}
	return nil
}