package userInfrastructure

import (
	"api.ruxwez.dev/adapters"
	"api.ruxwez.dev/adapters/mysql/models"
)

func FindUserByUsername(username string) *models.User {
	var user models.User

	if err := adapters.MySQL.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}

	return &user
}
