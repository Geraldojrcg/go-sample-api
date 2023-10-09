package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateRoutes(app *fiber.App, db *gorm.DB) {
	NewUserRoute(app, db)
	NewTodoRoute(app, db)
}
