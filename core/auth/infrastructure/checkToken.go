package authInfrastructure

import (
	"strconv"

	"internal/adapters"
	"internal/adapters/mysql/models"
	jwtLib "internal/libs/jwt"
)

func CheckToken(token string) int {
	// Cacheamos validaciones de tokens si es necesario (redis, memoria, etc.)
	data, err := jwtLib.Validate(token)
	if err != nil {
		return 0
	}

	// Evitamos castings innecesarios
	idStr, ok := data["id"].(string)
	if !ok {
		return 0
	}

	// Parseo de string a entero
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}

	// Consulta más eficiente a la base de datos usando SELECT COUNT en vez de cargar la estructura completa
	var count int64
	err = adapters.MySQL.Model(&models.AuthToken{}).Where("token = ?", token).Count(&count).Error
	if err != nil || count == 0 {
		return 0
	}

	return id
}
