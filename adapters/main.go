package adapters

import (
	"fmt"
	"log"

	"internal/adapters/mysql"
	"internal/adapters/mysql/models"
	"internal/vars"

	"gorm.io/gorm"
)

var (
	MySQL *gorm.DB
)

func Init() {
	log.Println("Abriendo conexión a los adaptadores")
	defer autoMigrate()
	// Inicializamos la base de datos MySQL principal
	var dbDsnPrimary = fmt.Sprintf("%s/database", vars.MYSQL_DSN) // Formulamos la línea DSN correspondiente
	// Inicializamos una conexión de MySQL
	MySQL = mysql.NewMySQLConnection(dbDsnPrimary)
}

func autoMigrate() {
	err := MySQL.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{}, &models.AuthToken{})
	if err != nil {
		log.Fatalln("Hubo un error al migrar la base de datos 'database' de MySQL")
	}
}
