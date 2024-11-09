package authApplication

import (
	"api.ruxwez.dev/adapters/mysql/models"
	sharedDomain "api.ruxwez.dev/core/_shared"
	authDomain "api.ruxwez.dev/core/auth/domain"
	authInfrastructure "api.ruxwez.dev/core/auth/infrastructure"
	userDomain "api.ruxwez.dev/core/user/domain"
	userInfrastructure "api.ruxwez.dev/core/user/infrastructure"
	passwordLib "api.ruxwez.dev/libs/password"
)

func ClassicUserRegister(http_body authDomain.ClassicRegisterBody) (string, error) {

	// Comprobamos si existe el nombre de usuario
	user := userInfrastructure.FindUserByUsername(http_body.Username)
	if user != nil {
		return "", authDomain.ErrUsernameAlreadyExists
	}

	// Comprobamos si existe el email
	user = userInfrastructure.FindUserByEmail(http_body.Email)
	if user != nil {
		return "", authDomain.ErrEmailAlreadyExists
	}

	// Encriptamos la contraseña
	encryptPass, err := passwordLib.Encrypt(http_body.Password)
	if err != nil {
		return "", sharedDomain.ErrInternalServer
	}

	// Creamos el modelo de usuario
	userModel := models.User{
		Username:  http_body.Username,
		Email:     http_body.Email,
		Password:  encryptPass,
		FirstName: http_body.FirstName,
		LastName:  http_body.LastName,
	}

	// Creamos el usuario en la base de datos
	err = userInfrastructure.CreateUser(&userModel)
	if err != nil {
		return "", userDomain.ErrToCreateUser
	}

	token := authInfrastructure.CreateAndRegisterAuthToken(userModel.ID)

	if token == "" {
		return "", sharedDomain.ErrInternalServer
	}

	return token, nil
}
