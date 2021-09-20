package Users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id             uint           `json:"id" gorm:"primary-key" `
	NIK            string         `json:"nik" gorm:"unique"`
	NamaDepan      string         `json:"namaDepan"`
	NamaBelakang   string         `json:"namaBelakang"`
	Email          string         `json:"email" gorm:"unique"`
	NoTelepon      string         `json:"noTelepon"`
	Alamat         string         `json:"alamat"`
	Kabupaten_Kota string         `json:"kabupaten_kota"`
	Provinsi       string         `json:"provinsi"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
