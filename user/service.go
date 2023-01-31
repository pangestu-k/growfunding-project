package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(requestInput RegisterUserInput) (User, error)
	LoginUser(requestInput LoginUserInput) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(requestInput RegisterUserInput) (User, error) {
	user := User{}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(requestInput.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.Name = requestInput.Name
	user.Email = requestInput.Email
	user.Occupation = requestInput.Occupation
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) LoginUser(requestInput LoginUserInput) (User, error) {
	email := requestInput.Email
	password := requestInput.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("Data with that email is not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}
