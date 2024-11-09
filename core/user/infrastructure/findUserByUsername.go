package userInfrastructure

import (
	"internal/adapters"
	"internal/adapters/mysql/models"
)

func FindUserByUsername(username string) *models.User {
	var user models.User

	if err := adapters.MySQL.Where("username = ?", username).First(&user).Error; err != nil {
		return nil
	}

	return &user
}
