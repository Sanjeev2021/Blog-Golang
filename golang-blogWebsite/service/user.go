package service

import (
	"blog/model"
	"errors"
	"time"
)

var allUsers []*model.User

func NewUser(id uint, name string, email string, password string) (*model.User, error) {

	if id < 0 {
		return nil, errors.New("The id cant be less than 0.")
	}

	if name == "" {
		return nil, errors.New("The name cant be empty")
	}

	if email == "" {
		return nil, errors.New("The email cant be empty")
	}

	if password == "" {
		return nil, errors.New("The password cant be empty")
	}

	createdAt := time.Now().Truncate(24 * time.Hour)
	updatedAt := time.Now().Truncate(24 * time.Hour)

	var tempUser = &model.User{
		ID:        id,
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	allUsers = append(allUsers, tempUser)
	return tempUser, nil

}

func UpdateUsername(u *model.User, name string) error {
	if u.Name == "" {
		return errors.New("the name cant be empty")
	}

	if name == "" {
		return errors.New("The name you are giving cant be empty.")
	}

	u.Name = name
	return nil
}

func DeleteUser(u *model.User) error {
	if u.Name == "" {
		return errors.New("The name cant be empty.")
	}

	u.Name = ""
	return nil
}

func GetUser() ([]*model.User, error) {
	return allUsers, nil
}
