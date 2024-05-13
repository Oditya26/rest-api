package models

type Buku struct {
	Id               int64  `gorm:"primaryKey" json:"id"`
	Judul            string `gorm:"type:varchar(255)" json:"judul" binding:"required"`
	Pengarang        string `gorm:"type:varchar(255)" json:"pengarang" binding:"required"`
	TanggalPublikasi string `gorm:"type:date" json:"tanggal_publikasi" binding:"required"`
	UpdatedAt        string `gorm:"type:date" json:"updated_at"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Buku) TableName() string {
	return "Buku"
}
