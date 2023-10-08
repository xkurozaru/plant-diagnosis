package database

import (
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
	"gorm.io/gorm"
)

type userDatabase struct {
	db *gorm.DB
}

func NewUserDatabase(db *gorm.DB) repository.UserRepository {
	return &userDatabase{db: db}
}

func (u userDatabase) Create(user model.User) error {
	userE := NewUserEntity(user)
	err := u.db.Create(&userE).Error
	if err != nil {
		return err
	}

	return nil
}

func (u userDatabase) Find(ID model.ULID) (model.User, error) {
	userE := UserEntity{}
	err := u.db.First(&userE, "id = ?", ID).Error
	if err != nil {
		return model.User{}, err
	}

	return userE.ToModel(), nil
}

func (u userDatabase) FindByLoginID(loginID string) (model.User, error) {
	userE := UserEntity{}
	err := u.db.First(&userE, "login_id = ?", loginID).Error
	if err != nil {
		return model.User{}, err
	}

	return userE.ToModel(), nil
}

func (u userDatabase) ExistsByLoginID(loginID string) (bool, error) {
	userEs := []UserEntity{}

	err := u.db.Where("login_id = ?", loginID).Find(&userEs).Error
	if err != nil {
		return false, err
	}

	return len(userEs) > 0, nil
}
