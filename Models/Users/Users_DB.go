package Users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id             uint           `gorm:"primaryKey; unique; not null" json:"id"`
	NIK            string         `gorm:"size:255; not null; unique" json:"nik"`
	NamaDepan      string         `gorm:"size:255; not null" json:"namaDepan"`
	NamaBelakang   string         `gorm:"size:255;not null" json:"namaBelakang"`
	TanggalLahir   string         `gorm:"type:date;not null" json:"tanggalLahir"`
	NoTelepon      string         `gorm:"size:255;not null" json:"noTelepon"`
	Alamat         string         `gorm:"size:1000;not null" json:"alamat"`
	Kabupaten_Kota string         `gorm:"size:255;not null" json:"kabupaten_kota"`
	Provinsi       string         `gorm:"size:255;not null" json:"provinsi"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
