package nilai

import (
	"context"
	"fmt"
	"log"
	"time"
	"tugas15/config"
	"tugas15/data"
	"tugas15/form"
)

const (
	table          = "nilai"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]data.Nilai, error) {
	var nilaiMhs []data.Nilai
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By id DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai data.Nilai
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Indeks,
			&nilai.Nilai,
			&nilai.MahasiswaID,
			&nilai.MataKuliahID,
			&createdAt,
			&updatedAt); err != nil {
				log.Println(nilai)
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

		log.Println(nilai)
	}
	return nilaiMhs, nil
}

func IsDuplicateNilai(ctx context.Context, mahasiswa_id, mata_kuliah_id int) (bool, error) {
    db, err := config.MySQL()
    if err != nil {
        return false, err
    }

    queryText := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE mahasiswa_id = %v AND mata_kuliah_id = %v", table, mahasiswa_id, mata_kuliah_id)
    var count int
    err = db.QueryRowContext(ctx, queryText).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func Insert(ctx context.Context, nilai form.InsertNilai) error {
	db, err := config.MySQL()
	if err != nil {
	  log.Fatal("Can't connect to MySQL", err)
	}
	
	queryText := fmt.Sprintf("INSERT IGNORE INTO %v (indeks, nilai, mata_kuliah_id, mahasiswa_id, created_at, updated_at) values('%v',%v,%v,%v, NOW(), NOW())", table,
	  	nilai.Indeks,
	  	nilai.Nilai,
		nilai.MataKuliahID,
		nilai.MahasiswaID,)
	_, err = db.ExecContext(ctx, queryText)
  
	if err != nil {
	  return err
	}
	return nil
}

func UpdateByID(ctx context.Context, nilai form.UpdateNilai, id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET indeks = '%v', nilai = %v, mata_kuliah_id = %v, mahasiswa_id = %v, updated_at = NOW() WHERE id = %v",
		table,
		nilai.Indeks,
		nilai.Nilai,
		nilai.MataKuliahID,
		nilai.MahasiswaID,
		id)
	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}

func DeleteByID(ctx context.Context, id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v WHERE id = %v", table, id)

	_, err = db.ExecContext(ctx, queryText)
	if err != nil {
		return err
	}

	return nil
}