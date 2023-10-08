package repository

import "github.com/xkurozaru/plant-diagnosis/controller/domain/model"

type UserRepository interface {
	Create(user model.User) error
	Find(ID model.ULID) (model.User, error)
	FindByLoginID(loginID string) (model.User, error)
	ExistsByLoginID(loginID string) (bool, error)
}
