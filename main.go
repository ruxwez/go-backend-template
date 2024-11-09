package main

import (
	"log"

	"internal/adapters"
	"internal/middlewares"
	"internal/routes"
	"internal/vars"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
}

func main() {
	vars.Init()     // Inicializamos las variables
	adapters.Init() // Inicializamos los adaptadores

	app := fiber.New() // Creamos una nueva instancia de Fiber

	middlewares.Init(app) // Inicializamos los middlewares

	routes.Init(app) // Inicializamos las rutas

	log.Fatalln(app.Listen(":" + vars.SERVER_PORT)) // Inicializamos el servidor
}
