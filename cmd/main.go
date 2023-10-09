package main

import (
	"github.com/geraldojrcg/go-sample-api/src/database"
	"github.com/geraldojrcg/go-sample-api/src/route"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	db := database.GetDatabaseConnection()

	route.CreateRoutes(app, db)

	log.Fatal(app.Listen(":3000"))
}
