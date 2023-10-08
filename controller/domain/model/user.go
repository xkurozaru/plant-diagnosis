package model

import "fmt"

const (
	minLoginIDLength = 4
	maxLoginIDLength = 32
)

type User struct {
	ID      ULID
	LoginID string
	Name    string
	Role    Role
}

func NewUser(loginID string, name string) (User, error) {
	u := User{
		ID:      NewULID(),
		LoginID: loginID,
		Name:    name,
		Role:    UserRole,
	}

	err := u.Validate()
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func NewAdminUser(loginID string, name string) (User, error) {
	u := User{
		ID:      NewULID(),
		LoginID: loginID,
		Name:    name,
		Role:    AdminRole,
	}

	err := u.Validate()
	if err != nil {
		return User{}, err
	}

	return u, nil
}

func (u User) Validate() error {
	if u.ID.IsEmpty() {
		return fmt.Errorf("ID must not be empty")
	}

	if len(u.LoginID) < minLoginIDLength || len(u.LoginID) > maxLoginIDLength {
		return fmt.Errorf("login ID length must be between %d and %d", minLoginIDLength, maxLoginIDLength)
	}

	if len(u.Name) <= 0 {
		return fmt.Errorf("name must not be empty")
	}

	return nil
}
