package services

import (
	"github.com/smugg99/coding-night-2023/models"
	"gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

func GetUserByCredentials(db *gorm.DB, creds *models.UserLoginDTO) (authorized bool, user models.User) {
    if err := db.Find(&user, "email = ?", creds.Email).Error; err != nil {
        return false, models.User{}
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err == nil {
        return true, user
    }

    return false, models.User{}
}

func CreateNewUser(db *gorm.DB, newUser *models.User) error {
    return db.Create(newUser).Error
}

