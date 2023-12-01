package models

import (
	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email string `json:"email" validate:"email,required"`
    Password string `json:",omitempty" validate:"bcrypt,required"`

    // optional (contact info)
    Name string `json:"name"`
    Phone string `json:"phone_number" validate:"e164"`
    Address string `json:"address"`
}

type UserRegisterDTO struct {
    Email string `json:"email" validate:"email,required"`
    Password string `json:"password" validate:"required"`

    // optional (contact info)
    Name string `json:"name"`
    Phone string `json:"phone_number" validate:"e164"`
    Address string `json:"address"`

}

type UserLoginDTO struct {
    Email string `json:"email" validate:"email,required"`
    Password string `json:"password" validate:"required"`
}

