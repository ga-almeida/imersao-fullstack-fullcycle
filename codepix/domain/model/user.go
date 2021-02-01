package model

import (
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Base     `valid:"required"`
	Name     string     `json:"name" valid:"notnull"`
	Email    string     `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func NewUser(email string, name string) (*User, error) {
	user := User{
		Email: email,
		Name: name,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
