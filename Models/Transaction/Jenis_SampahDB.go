package Transaction

import (
	"time"

	"gorm.io/gorm"
)

type JenisSampah struct {
	Id         uint           `gorm:"primaryKey; unique; not null" json:"id"`
	NamaJenis  string         `gorm:"size:255; not null" json:"namaJenis"`
	MinimalQty int            `gorm:"size:10; not null" json:"minimalQty"`
	Satuan     string         `gorm:"size:255; not null" json:"satuan"`
	HargaJual  float32        `gorm:"type:decimal(13,2)" json:"hargaJual"`
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
