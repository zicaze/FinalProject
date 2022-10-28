package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username     string    `gorm:"not null;uniqueIndex" json:"username" valid:"required"`
	Email        string    `gorm:"not null;uniqueIndex" json:"email" valid:"email,required"`
	Password     string    `gorm:"not null" json:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minumum length of 6 characters"`
	Age          uint      `gorm:"not null" json:"age" valid:"numeric, range(9|100),required"`
	Photos       []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Comments     []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SocialMedias []Sosmed  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type UpdateUserInput struct {
	Username string `json:"username" gorm:"not null" valid:"required"`
	Email    string `json:"email" gorm:"not null" valid:"email,required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}

func (u *UpdateUserInput) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	err = nil
	return
}
