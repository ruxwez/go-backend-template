package authApplication

import (
	sharedDomain "api.ruxwez.dev/core/_shared"
	authDomain "api.ruxwez.dev/core/auth/domain"
	authInfrastructure "api.ruxwez.dev/core/auth/infrastructure"
	userInfrastructure "api.ruxwez.dev/core/user/infrastructure"
	passwordLib "api.ruxwez.dev/libs/password"
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
