package OperatorSampahDriver

import (
	"errors"
	"golang-final-project/Models/BankSampah"

	"gorm.io/gorm"
)

// Register Operator Sampah
func CreateOperatorSampah(db *gorm.DB, os *BankSampah.OperatorSampah) error {
	if err := db.Create(&os).Error; err != nil {
		return err
	}
	return nil
}

func GetAllOperatorSampah(db *gorm.DB) ([]BankSampah.OperatorSampah, error) {
	operatorSampah := []BankSampah.OperatorSampah{}

	if err := db.Preload("BankSampah").Find(&operatorSampah).Error; err != nil {
		return operatorSampah, err
	}

	return operatorSampah, nil
}

func GetOperatorSampahByID(id string, db *gorm.DB) (BankSampah.OperatorSampah, bool, error) {
	var operatorSampah BankSampah.OperatorSampah

	err := db.Preload("BankSampah").First(&operatorSampah, id).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return operatorSampah, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return operatorSampah, false, nil
	}
	return operatorSampah, true, nil
}

func UpdateOperatorSampah(id string, db *gorm.DB, op *BankSampah.OperatorSampah) error {
	if err := db.Model(BankSampah.OperatorSampah{}).Where("id = ?", id).Updates(op).Error; err != nil {
		return err
	}
	return nil
}

func DeleteOperatorSampah(id string, db *gorm.DB) error {
	var operatorSampah BankSampah.OperatorSampah
	if err := db.Where("id = ? ", id).Unscoped().Delete(&operatorSampah).Error; err != nil {
		return err
	}
	return nil
}

func GetOperatorSampahByObject(db *gorm.DB, op *BankSampah.OperatorSampah) (BankSampah.OperatorSampah, error) {
	if err := db.Preload("BankSampah").Find(&op).Error; err != nil {
		return *op, err
	}

	return *op, nil
}
