package Login

import (
	"golang-final-project/Models/Users"
	"time"

	"gorm.io/gorm"
)

type LoginDataUsers struct {
	UserId    uint           `gorm:"not null; PrimaryKey" json:"userId"`
	Email     string         `gorm:"size:255; not null; unique" json:"email"`
	Username  string         `gorm:"size:100; not null; unique" json:"username"`
	Password  string         `gorm:"size:100; not null" json:"password"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	User      Users.User     `gorm:"ForeignKey:UserId;References:Id;constraint:OnDelete:CASCADE;"`
}

type Login struct {
	Email    string `form:"email"`
	Username string `form:"username"`
	Password string `form:"password"`
}
