package services

import (
	"my-orm/database"
	"my-orm/models"
	"my-orm/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
}

func (us *UserService) Register(userReq models.UserRequest) (models.User, error) {

	password, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.User{}, err

	}
	var user models.User = models.User{

		Email:    userReq.Email,
		Password: string(password),
	}

	result := database.DB.Create(&user)

	if err := result.Error; err != nil {
		return models.User{}, err
	}

	err = result.Last(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (us *UserService) Login(userReq models.UserRequest) (string, error) {
	var user models.User
	err := database.DB.First(&user, "email = ?", userReq.Email).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userReq.Password))

	if err != nil {
		return "", err
	}

	token, err := utils.GenerateJWT(int(user.ID), models.JWTOptions{
		SecretKey:      "secret",
		ExpireDuration: 1,
	})
	if err != nil {
		return "", err
	}

	return token, nil

}
