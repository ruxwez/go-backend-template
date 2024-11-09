package usernameLib

import "regexp"

func CheckFormat(username string) bool {
	// Expresión regular que permite letras, números, puntos, guiones y guiones bajos, y entre 3 y 12 caracteres
	re := regexp.MustCompile(`^[a-zA-Z0-9._-]{3,12}$`)
	return re.MatchString(username)
}
