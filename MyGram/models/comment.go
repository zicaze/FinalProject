package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	CModel
	Message  string `gorm:"not null" json:"message" valid:"required"`
	Photo_id string `json:"photo_id"`
	User_id  string `json:"user_id"`
	User     *Comments
	Photo    *Photos
}

type Comments struct {
	GormModel
	Username string `json:"username"`
	Email    string `json:"email"`
}

type Photos struct {
	CModel
	Title     string `json:"title"`
	Caption   string `json:"caption"`
	Photo_url string `json:"photo_url"`
	User_id   uint   `json:"user_id"`
}

type FormComment struct {
	Message string `gorm:"not null" json:"message" valid:"required"`
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	err = nil
	return
}

func (u *FormComment) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errcreate := govalidator.ValidateStruct(u)

	if errcreate != nil {
		err = errcreate
		return
	}
	err = nil
	return
}

func (p *Comment) GetComments(db *gorm.DB) (*[]Comment, error) {
	var err error
	Comments := []Comment{}
	err = db.Debug().Model(&Comment{}).Limit(100).Find(&Comments).Error
	if err != nil {
		return &[]Comment{}, err
	}
	if len(Comments) > 0 {
		for i := range Comments {
			er := db.Debug().Select("id", "username", "email").Model(&User{}).Where("id = ?", Comments[i].User_id).Find(&Comments[i].User).Error
			if er != nil {
				return &[]Comment{}, er
			}
			err := db.Joins("photos", db.Select("id", "title", "caption", "photo_url", "user_id").Where("id = ?", Comments[i].Photo_id).Model(&Photo{}).Find(&Comments[i].Photo)).Error
			if err != nil {
				return &[]Comment{}, err
			}
		}
	}
	return &Comments, nil
}
