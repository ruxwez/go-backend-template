package jwtLibGen

import jwtLib "api.ruxwez.dev/libs/jwt"

func GenerateAuthToken(userID string) (string, error) {
	// Instanciamos el map para añadir el user id al claims
	data := make(map[string]interface{})
	data["id"] = userID
	data["type"] = "authToken"

	// Generamos el token
	token, err := jwtLib.Generate(data)
	if err != nil {
		return "", err
	}

	return token, nil
}
