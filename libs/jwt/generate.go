package jwtLib

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

func Generate(claimsMap map[string]interface{}) (string, error) {
	// Instanciamos los claims y definimos la expiracion, proximamente dinámica
	claims := jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 168).Unix(), // Expira en 7 días
	}

	// Añadimos los diferentes claims dinámicos
	for key, value := range claimsMap {
		claims[key] = value
	}

	// Generamos el token y lo firmamos
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
