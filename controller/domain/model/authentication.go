package model

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8
	maxPasswordLength = 72
)

type Authentication struct {
	ID             ULID
	User           User
	HashedPassword string
}

func NewAuthentication(user User, password string) (Authentication, error) {
	if len(password) < minPasswordLength || len(password) > maxPasswordLength {
		return Authentication{}, fmt.Errorf("password length must be between %d and %d", minPasswordLength, maxPasswordLength)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Authentication{}, err
	}

	a := Authentication{
		ID:             NewULID(),
		User:           user,
		HashedPassword: string(hash),
	}

	err = a.Validate()
	if err != nil {
		return Authentication{}, err
	}

	return a, nil
}

func (a Authentication) Validate() error {
	if a.ID.IsEmpty() {
		return fmt.Errorf("ID must not be empty")
	}

	if a.User.ID.IsEmpty() {
		return fmt.Errorf("user ID must not be empty")
	}

	if len(a.HashedPassword) != 60 {
		return fmt.Errorf("hash length must be 60")
	}

	return nil
}

func (a Authentication) Authenticate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(a.HashedPassword), []byte(password))
}
