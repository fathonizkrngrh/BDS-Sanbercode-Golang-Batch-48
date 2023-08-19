package form

type (
	InsertMahasiswa struct {
		Nama string `json:"nama" binding:"required"`
	}

	UpdateMahasiswa struct {
		Nama string `json:"Nama" binding:"required"`
	}
)