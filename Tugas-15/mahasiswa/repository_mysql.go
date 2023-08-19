package mahasiswa

import (
	"context"
	"fmt"
	"log"
	"time"
	"tugas15/config"
	"tugas15/data"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAll(ctx context.Context) ([]data.Mahasiswa, error) {
	var mahasiswa []data.Mahasiswa
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
		var mhs data.Mahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&mhs.ID,
			&mhs.Nama,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		mhs.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		mhs.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		mahasiswa = append(mahasiswa, mhs)
	}
	return mahasiswa, nil
}

func IsDuplicateMahasiswa(ctx context.Context, nama string) (bool, error) {
    db, err := config.MySQL()
    if err != nil {
        return false, err
    }

    queryText := fmt.Sprintf("SELECT COUNT(*) FROM %v WHERE nama = '%v'", table, nama)
    var count int
    err = db.QueryRowContext(ctx, queryText).Scan(&count)
    if err != nil {
        return false, err
    }

    return count > 0, nil
}

func Insert(ctx context.Context, mahasiswa data.Mahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
	  log.Fatal("Can't connect to MySQL", err)
	}
	
	queryText := fmt.Sprintf("INSERT IGNORE INTO %v (nama, created_at, updated_at) values('%v', NOW(), NOW())", table,
	  	mahasiswa.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
	  return err
	}

	return nil
}

func UpdateByID(ctx context.Context, mahasiswa data.Mahasiswa, id int) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET nama = '%v', updated_at = NOW() WHERE id = %v",
		table,
		mahasiswa.Nama,
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