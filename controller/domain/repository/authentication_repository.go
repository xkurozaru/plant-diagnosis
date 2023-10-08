package repository

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type AuthenticationRepository interface {
	Create(authentication model.Authentication) error
	FindByUserID(userID model.ULID) (model.Authentication, error)
}
