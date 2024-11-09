package userInfrastructure

import (
	"log"

	"internal/adapters"
	"internal/adapters/mysql/models"
	userDomain "internal/core/user/domain"
)

func CreateUser(user *models.User) error {
	if err := adapters.MySQL.Create(user).Error; err != nil {
		log.Println(err) // Logeamos el error
		return userDomain.ErrToCreateUser
	}

	return nil
}
