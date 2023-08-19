package form

type (
	InsertNilai struct {
		Nilai        int    `json:"nilai" binding:"required"`
		Indeks       string `json:"indeks" binding:"required"`
		MataKuliahID int    `json:"mata_kuliah_id" binding:"required"`
		MahasiswaID  int    `json:"mahasiswa_id" binding:"required"`
	}

	UpdateNilai struct {
		Nilai        int    `json:"nilai" binding:"required"`
		Indeks       string `json:"indeks" binding:"required"`
		MataKuliahID int    `json:"mata_kuliah_id" binding:"required"`
		MahasiswaID  int    `json:"mahasiswa_id" binding:"required"`
	}
)