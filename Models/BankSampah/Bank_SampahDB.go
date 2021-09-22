package BankSampah

import (
	"time"

	"gorm.io/gorm"
)

type BankSampah struct {
	Id             int            `gorm:"primaryKey; auto_increment; not null" json:"id"`
	NamaUsaha      string         `gorm:"size:255; not null" json:"namaUsaha"`
	NamaPemilik    string         `gorm:"size:255; not null" json:"namaPemilik"`
	NIB            string         `gorm:"size:255; not null; unique" json:"nib"`
	NoTelepon      string         `gorm:"size:255; not null; unique" json:"noTelepon"`
	EmailResmi     string         `gorm:"size:255; not null; unique" json:"emailResmi"`
	Alamat         string         `gorm:"size:1000; not null" json:"alamat"`
	Kabupaten_Kota string         `gorm:"size:255; not null" json:"kabupaten_kota"`
	Provinsi       string         `gorm:"size:255; not null" json:"provinsi"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
