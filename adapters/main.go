package adapters

import (
	"fmt"
	"log"

	"api.ruxwez.dev/adapters/mysql"
	"api.ruxwez.dev/vars"
	"gorm.io/gorm"
)

var (
	MySQL *gorm.DB
)

func Init() {
	log.Println("Abriendo conexión a los adaptadores")
	defer autoMigrate()
	// Inicializamos la base de datos MySQL principal
	var dbDsnPrimary = fmt.Sprintf("%s/ruxwez", vars.MYSQL_DSN) // Formulamos la línea DSN correspondiente
	// Inicializamos una conexión de MySQL
	MySQL = mysql.NewMySQLConnection(dbDsnPrimary)
}

func autoMigrate() {
	err := MySQL.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate()
	if err != nil {
		log.Fatalln("Hubo un error al migrar la base de datos 'ruxwez' de MySQL")
	}
}
