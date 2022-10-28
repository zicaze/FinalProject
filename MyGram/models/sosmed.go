package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Sosmed struct {
	SModel
	Name             string `gorm:"not null" json:"name" valid:"required"`
	Social_media_url string `gorm:"not null" json:"social_media_url" valid:"required"`
	User_id          uint   `json:"user_id"`
	User             *Userso
}

type Userso struct {
	GormModel
	Username  string `json:"username"`
	Photo_url string `json:"profile_image_url"`
}

type FormSosmed struct {
	GormModel
	Name             string `json:"name" valid:"required"`
	Social_media_url string `json:"social_media_url" valid:"required"`
}

func (u *Sosmed) BeforeCreate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	err = nil
	return
}

func (u *FormSosmed) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	err = nil
	return
}

func (p *Sosmed) GetSosmeds(db *gorm.DB) (*[]Sosmed, error) {
	var err error
	Sosmeds := []Sosmed{}
	err = db.Debug().Model(&Sosmed{}).Limit(100).Find(&Sosmeds).Error
	if err != nil {
		return &[]Sosmed{}, err
	}
	if len(Sosmeds) > 0 {
		for i := range Sosmeds {
			err := db.Joins("users", db.Select("id", "username").Where("id = ?", Sosmeds[i].User_id).Model(&User{}).Find(&Sosmeds[i].User)).Error
			if err != nil {
				return &[]Sosmed{}, err
			}
			er := db.Joins("photos", db.Select("photo_url").Where("user_id= ?", Sosmeds[i].User_id).Model(&Photo{}).Find(&Sosmeds[i].User)).Error
			if er != nil {
				return &[]Sosmed{}, er
			}
		}
	}
	return &Sosmeds, nil
}
