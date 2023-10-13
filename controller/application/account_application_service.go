package application

import (
	"fmt"

	"github.com/xkurozaru/plant-diagnosis/controller/domain/model"
	"github.com/xkurozaru/plant-diagnosis/controller/domain/repository"
)

type AccountApplicationService interface {
	SignUp(name string, loginID string, password string) error
	SignUpAdmin(name string, loginID string, password string) error
	SignIn(loginID string, password string) (model.User, error)
	GetUser(userID model.ULID) (model.User, error)
}

type accountApplicationService struct {
	authenticationRepository repository.AuthenticationRepository
	userRepository           repository.UserRepository
}

func NewAccountApplicationService(
	authenticationRepository repository.AuthenticationRepository,
	userRepository repository.UserRepository,
) AccountApplicationService {
	return accountApplicationService{
		authenticationRepository: authenticationRepository,
		userRepository:           userRepository,
	}
}

func (a accountApplicationService) SignUp(name string, loginID string, password string) error {
	user, err := model.NewUser(loginID, name)
	if err != nil {
		return err
	}

	auth, err := model.NewAuthentication(user, password)
	if err != nil {
		return err
	}

	exists, err := a.userRepository.ExistsByLoginID(loginID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("this loginID user already exists")
	}

	err = a.userRepository.Create(user)
	if err != nil {
		return err
	}

	err = a.authenticationRepository.Create(auth)
	if err != nil {
		return err
	}

	return nil
}

func (a accountApplicationService) SignUpAdmin(name string, loginID string, password string) error {
	user, err := model.NewAdminUser(loginID, name)
	if err != nil {
		return err
	}

	auth, err := model.NewAuthentication(user, password)
	if err != nil {
		return err
	}

	exists, err := a.userRepository.ExistsByLoginID(loginID)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("this loginID user already exists")
	}

	err = a.userRepository.Create(user)
	if err != nil {
		return err
	}

	err = a.authenticationRepository.Create(auth)
	if err != nil {
		return err
	}

	return nil
}

func (a accountApplicationService) SignIn(loginID string, password string) (model.User, error) {
	user, err := a.userRepository.FindByLoginID(loginID)
	if err != nil {
		return model.User{}, err
	}

	auth, err := a.authenticationRepository.FindByUserID(user.ID)
	if err != nil {
		return model.User{}, err
	}

	err = auth.Authenticate(password)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (a accountApplicationService) GetUser(userID model.ULID) (model.User, error) {
	user, err := a.userRepository.Find(userID)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
