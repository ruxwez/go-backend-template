package vars

import (
	"log"
	"os"
)

var (
	SMTP_HOST     string
	SMTP_PORT     string
	SMTP_EMAIL    string
	SMTP_PASSWORD string
)

func initSMTP() {
	SMTP_HOST = os.Getenv("SMTP_HOST")
	if SMTP_HOST == "" {
		log.Fatalln("No se ha definido la variable de entorno 'SMTP_HOST'")
	}

	SMTP_PORT = os.Getenv("SMTP_PORT")
	if SMTP_PORT == "" {
		SMTP_PORT = "25"
	}

	SMTP_EMAIL = os.Getenv("SMTP_EMAIL")
	if SMTP_EMAIL == "" {
		log.Fatalln("No se ha definido la variable de entorno 'SMTP_EMAIL'")
	}

	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")
	if SMTP_PASSWORD == "" {
		log.Fatalln("No se ha definido la variable de entorno 'SMTP_PASSWORD'")
	}
}
