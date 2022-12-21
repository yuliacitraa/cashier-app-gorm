package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db}
}

func (u *UserRepository) AddUser(user model.User) error {
	err := u.db.Create(&user).Error
	return err
}

func (u *UserRepository) UserAvail(cred model.User) error {
	// ambil data dari db
	var usr model.User 
	err := u.db.Where("username = ?", cred.Username).Find(&usr).Error
	if err != nil {
		return err
	}

	if usr.Username == "" {
		return errors.New("not found") //make new message error
	}

	if usr.Username == cred.Username {
		if usr.Password != cred.Password {
			return errors.New("wrong password")
		}
	}
	return nil
}

func (u *UserRepository) CheckPassLength(pass string) bool {
	if len(pass) <= 5 {
		return true
	}

	return false
}

func (u *UserRepository) CheckPassAlphabet(pass string) bool {
	for _, charVariable := range pass {
		if (charVariable < 'a' || charVariable > 'z') && (charVariable < 'A' || charVariable > 'Z') {
			return false
		}
	}
	return true
}
