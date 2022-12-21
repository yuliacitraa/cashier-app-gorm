package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	err := u.db.Create(&session).Error
	return err
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	err := u.db.Where("token = ?", tokenTarget).Delete(&model.Session{}).Error
	return err
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	err := u.db.Model(&session).Where("username = ?", session.Username).Update("token", session.Token).Update("expiry", session.Expiry).Error
	return err
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session := model.Session{}
	err := u.db.Where("token = ?", token).First(&session).Error
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSessions(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return session, nil
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	var session model.Session 
	err := u.db.Where("username = ?", name).First(&session).Error
	if err != nil {
		return model.Session{}, err
	}

	return session, nil}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session 
	err := u.db.Where("token = ?", token).First(&session).Error
	if err != nil {
		return model.Session{}, err
	}
	return session, nil}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
