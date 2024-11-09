package userInfrastructure

import (
	"log"

	"api.ruxwez.dev/adapters"
	"api.ruxwez.dev/adapters/mysql/models"
	userDomain "api.ruxwez.dev/core/user/domain"
)

func CreateUser(user *models.User) error {
	if err := adapters.MySQL.Create(user).Error; err != nil {
		log.Println(err) // Logeamos el error
		return userDomain.ErrToCreateUser
	}

	return nil
}
