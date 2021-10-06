package JenisSampahDriver

import (
	"errors"
	"golang-final-project/Models/Transaction"

	"gorm.io/gorm"
)

func CreateJenisSampah(db *gorm.DB, js *Transaction.JenisSampah) error {
	if err := db.Create(&js).Error; err != nil {
		return err
	}
	return nil
}

func GetAllJenisSampah(db *gorm.DB) ([]Transaction.JenisSampah, error) {
	jenisSampah := []Transaction.JenisSampah{}

	if err := db.Find(&jenisSampah).Error; err != nil {
		return jenisSampah, err
	}

	return jenisSampah, nil
}

func GetJenisSampahByID(id string, db *gorm.DB) (Transaction.JenisSampah, bool, error) {
	js := Transaction.JenisSampah{}

	err := db.Where("id = ?", id).First(&js).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return js, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return js, false, nil
	}
	return js, true, nil
}

func UpdateJenisSampah(db *gorm.DB, js *Transaction.JenisSampah, id string) error {
	if err := db.Model(Transaction.JenisSampah{}).Where("id = ?", id).Updates(js).Error; err != nil {
		return err
	}
	return nil
}

func DeleteJenisSampah(id string, db *gorm.DB) error {
	var js Transaction.JenisSampah
	if err := db.Where("id = ? ", id).Unscoped().Delete(&js).Error; err != nil {
		return err
	}
	return nil
}
