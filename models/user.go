package models

import (
	"SchoolTest/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"not null" json:"username" valid:"required~Username is required!"`
	Password string `gorm:"not null" json:"password" valid:"required~Password is required!,minstringlength(6)~Password minimum length must be 6 characters!"`
	Role int `gorm:"not null" json:"role" valid:"required~Role is required!"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errC := govalidator.ValidateStruct(u)

	if errC != nil {
		err = errC
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}