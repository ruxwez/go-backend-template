package userInfrastructure

import (
	"internal/adapters"
	"internal/adapters/mysql/models"
)

func FindUserByEmail(email string) *models.User {
	var user models.User

	if err := adapters.MySQL.Where("email = ?", email).First(&user).Error; err != nil {
		return nil
	}

	return &user
}
