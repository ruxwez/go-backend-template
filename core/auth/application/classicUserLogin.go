package authApplication

import (
	sharedDomain "internal/core/_shared"
	authDomain "internal/core/auth/domain"
	authInfrastructure "internal/core/auth/infrastructure"
	userInfrastructure "internal/core/user/infrastructure"
	passwordLib "internal/libs/password"
)

func ClassicUserLogin(http_body authDomain.ClassicLoginBody) (string, error) {

	// Obtenemos el usuario por el email
	user := userInfrastructure.FindUserByEmail(http_body.Email)
	if user == nil {
		return "", authDomain.ErrInvalidCredentials
	}

	// Comprobamos si la contraseña es correcta
	if !passwordLib.Validate(http_body.Password, user.Password) {
		return "", authDomain.ErrInvalidCredentials
	}

	// Creamos el token de autenticación
	token := authInfrastructure.CreateAndRegisterAuthToken(user.ID)
	if token == "" {
		return "", sharedDomain.ErrInternalServer
	}

	return token, nil
}
