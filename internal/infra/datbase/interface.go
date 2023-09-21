package database

import "github.com/lucasferreirajs/fc-goexpert/apis/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

func (u *User) Create(user *entity.User) error {

	return u.DB.Create(user).Error

}

func (u *User) FindByEmail(email string) (*entity.User, error) {

	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
