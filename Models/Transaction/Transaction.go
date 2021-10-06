package Transaction

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	Id                uint           `gorm:"PrimaryKey; Auto Increment; Not Null" json:"id"`
	BankSampahId      uint           `gorm:"not null" json:"bankSampahId"`
	UserId            uint           `gorm:"not null; ForeignKey; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"userId"`
	OperatorSampahId  uint           `gorm:"not null" json:"operatorId"`
	TotalQty          int            `gorm:"size:10" json:"totalQty"`
	TotalTransaction  float32        `gorm:"type:decimal(13,2)" json:"totalTransaction"`
	Status            string         `gorm:"size:255;default:process" json:"status"`
	CreatedAt         time.Time      `json:"createdAt"`
	UpdatedAt         time.Time      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	DetailTransaction []DetailTransaction
}

type DetailTransaction struct {
	TransactionID   uint    `gorm:"Not Null; PrimaryKey; autoIncrement:false; ForeignKey; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"transactionId"`
	JenisSampahId   uint    `gorm:"Not Null; PrimaryKey; autoIncrement:false" json:"jenisSampahId"`
	Qty             int     `gorm:"size:10; not null" json:"qty"`
	TotalPrice      float32 `gorm:"type:decimal(13,2)" json:"totalPrice"`
	TransactionType string
}

type TransactionReq struct {
	BankSampahId     uint                `json:"bankSampahId"`
	UserId           uint                `json:"userId"`
	OperatorSampahId uint                `json:"operatorId"`
	TotalQty         int                 `json:"totalQty"`
	TotalTransaction float32             `json:"totalTransaction"`
	Status           string              `json:"status"`
	Detail           []DetailTransaction `json:"detail"`
}
