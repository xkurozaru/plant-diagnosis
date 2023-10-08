package database

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type AuthenticationEntity struct {
	ID             string `gorm:"primaryKey"`
	UserID         string
	User           UserEntity `gorm:"foreignKey:UserID"`
	HashedPassword string
}

func NewAuthenticationEntity(authentication model.Authentication) AuthenticationEntity {
	return AuthenticationEntity{
		ID:             authentication.ID.ToString(),
		UserID:         authentication.User.ID.ToString(),
		HashedPassword: authentication.HashedPassword,
	}
}

func (a AuthenticationEntity) ToModel() model.Authentication {
	return model.Authentication{
		ID:             model.ULID(a.ID),
		User:           a.User.ToModel(),
		HashedPassword: a.HashedPassword,
	}
}
