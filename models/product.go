package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	Title       string `gorm:"title" json:"title" valid:"required~The Title for the product is required"`
	Description string `gorm:"description" json:"description" valid:"required~Description for the product is required"`
	UserID      uint
	User        *User
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {

	_, errUpdate := govalidator.ValidateStruct(p)

	if errUpdate != nil {
		err = errUpdate
		return
	}

	err = nil
	return
}
