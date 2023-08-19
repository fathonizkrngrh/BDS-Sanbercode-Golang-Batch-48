package form

type (
	InsertMataKuliah struct {
		Nama string `json:"nama" binding:"required"`
	}

	UpdateMataKuliah struct {
		Nama string `json:"Nama" binding:"required"`
	}
)