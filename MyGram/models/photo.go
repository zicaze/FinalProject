package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	PModel
	Title     string `gorm:"not null" json:"title" valid:"required"`
	Caption   string `json:"caption"`
	Photo_url string `gorm:"not null" json:"photo_url" valid:"required"`
	User_id   uint
	User      *Users
}

type Users struct {
	PModel
	Username string `json:"username"`
	Email    string `json:"email"`
}

type FormPhoto struct {
	Title     string `gorm:"not null" json:"title" valid:"required"`
	Caption   string `json:"caption"`
	Photo_url string `gorm:"not null" json:"photo_url" valid:"required"`
}

func (u *Photo) BeforeCreate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	err = nil
	return
}

func (u *FormPhoto) BeforeCreate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	err = nil
	return
}

func (p *Photo) GetPhotos(db *gorm.DB) (*[]Photo, error) {
	var err error
	Photos := []Photo{}
	err = db.Debug().Model(&Photo{}).Limit(100).Find(&Photos).Error
	if err != nil {
		return &[]Photo{}, err
	}
	if len(Photos) > 0 {
		for i := range Photos {
			err := db.Debug().Select("username", "email").Where("id = ?", Photos[i].User_id).Take(&Photos[i].User).Error
			if err != nil {
				return &[]Photo{}, err
			}
		}
	}
	return &Photos, nil
}
