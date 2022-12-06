package main

import (
	"fmt"
	"log"
	"os"
	"support-ticket/database"
	"support-ticket/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	fastergoding.Run() // hot reload
	database.Connect()
	app := fiber.New()
	routes.Setup(app)

	errListen := app.Listen(":3002")
	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		fmt.Println(errListen)
		os.Exit(1)
	}
}
