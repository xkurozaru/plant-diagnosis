package database

import (
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
	"gorm.io/gorm"
)

type authenticationDatabase struct {
	db *gorm.DB
}

func NewAuthenticationDatabase(db *gorm.DB) repository.AuthenticationRepository {
	return &authenticationDatabase{db: db}
}

func (a authenticationDatabase) Create(authentication model.Authentication) error {
	authenticationE := NewAuthenticationEntity(authentication)

	err := a.db.Create(&authenticationE).Error
	if err != nil {
		return err
	}

	return nil
}

func (a authenticationDatabase) FindByUserID(userID model.ULID) (model.Authentication, error) {
	authenticationE := AuthenticationEntity{}
	err := a.db.Where("user_id = ?", userID.ToString()).First(&authenticationE).Error
	if err != nil {
		return model.Authentication{}, err
	}

	return authenticationE.ToModel(), nil
}
