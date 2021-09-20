package Users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id             uint           `json:"id" gorm:"primary-key" `
	NIK            string         `json:"nik" gorm:"unique;size:256"`
	NamaDepan      string         `json:"namaDepan" gorm:"size:256"`
	NamaBelakang   string         `json:"namaBelakang" gorm:"size:256"`
	Email          string         `json:"email" gorm:"unique"`
	TanggalLahir   string         `json:"tanggalLahir" gorm:"type:date"`
	NoTelepon      string         `json:"noTelepon" gorm:"size:256"`
	Alamat         string         `json:"alamat"`
	Kabupaten_Kota string         `json:"kabupaten_kota" gorm:"size:256"`
	Provinsi       string         `json:"provinsi" gorm:"size:256"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
