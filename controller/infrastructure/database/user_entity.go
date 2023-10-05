package database

import (
	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
)

type UserEntity struct {
	Model
	LoginID string
	Name    string
	Role    string
}

func NewUserEntity(u model.User) UserEntity {
	return UserEntity{
		Model:   Model{ID: u.ID.ToString()},
		LoginID: u.LoginID,
		Name:    u.Name,
		Role:    u.Role.Type,
	}
}

func (e UserEntity) ToModel() model.User {
	return model.User{
		ID:      model.ULID(e.ID),
		LoginID: e.LoginID,
		Name:    e.Name,
		Role:    model.NewRole(e.Role),
	}
}
