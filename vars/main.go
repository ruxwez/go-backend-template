package vars

import (
	"log"
	"os"
)

var (
	MYSQL_DSN   string
	SERVER_PORT string
)

func Init() {
	MYSQL_DSN = os.Getenv("MYSQL_DSN")
	if MYSQL_DSN == "" {
		log.Fatalln("No se ha definido la variable de entorno 'MYSQL_DSN'")
	}

	SERVER_PORT = os.Getenv("SERVER_PORT")
	if SERVER_PORT == "" {
		SERVER_PORT = "3000"
	}
}
