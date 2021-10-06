package BankSampahDriver

import (
	"errors"
	"golang-final-project/Models/BankSampah"

	"gorm.io/gorm"
)

func CreateBankSampah(db *gorm.DB, b *BankSampah.BankSampah) error {
	if err := db.Create(&b).Error; err != nil {
		return err
	}
	return nil
}

func GetAllBankSampah(db *gorm.DB) ([]BankSampah.BankSampah, error) {
	bankSampah := []BankSampah.BankSampah{}

	if err := db.Find(&bankSampah).Error; err != nil {
		return bankSampah, err
	}

	return bankSampah, nil
}

func GetBankSampahByID(id string, db *gorm.DB) (BankSampah.BankSampah, bool, error) {
	b := BankSampah.BankSampah{}

	err := db.Where("id = ?", id).First(&b).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return b, false, nil
	}
	return b, true, nil
}

func DeleteBankSampah(id string, db *gorm.DB) error {
	var b BankSampah.BankSampah
	if err := db.Where("id = ? ", id).Unscoped().Delete(&b).Error; err != nil {
		return err
	}
	return nil
}

func UpdateBankSampah(db *gorm.DB, b *BankSampah.BankSampah, id string) error {
	if err := db.Model(BankSampah.BankSampah{}).Where("id = ?", id).Updates(b).Error; err != nil {
		return err
	}
	return nil
}
