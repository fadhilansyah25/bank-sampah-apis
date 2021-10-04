package UserDriver

import (
	"errors"
	"golang-final-project/Models/Users"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, u *Users.User) error {
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(db *gorm.DB) ([]Users.User, error) {
	users := []Users.User{}

	if err := db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func GetUserByID(id string, db *gorm.DB) (Users.User, bool, error) {
	u := Users.User{}

	err := db.Where("id = ?", id).First(&u).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return u, false, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return u, false, nil
	}
	return u, true, nil
}

func DeleteUser(id string, db *gorm.DB) error {
	var u Users.User
	if err := db.Where("id = ? ", id).Unscoped().Delete(&u).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(db *gorm.DB, u *Users.User, id string) error {

	if err := db.Model(Users.User{}).Where("id = ?", id).Updates(u).Error; err != nil {
		return err
	}
	return nil
}
