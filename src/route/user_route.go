package route

import (
	"github.com/geraldojrcg/go-sample-api/src/controller"
	"github.com/geraldojrcg/go-sample-api/src/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserRouter struct {
	controller controller.UserController
}

func (u *UserRouter) Load(r *fiber.App) {
	r.Get("/users", u.controller.GetAll)
	r.Get("/users/:id", u.controller.GetById)
	r.Post("/users", u.controller.Create)
	r.Patch("/users/:id", u.controller.Update)
	r.Delete("/users/:id", u.controller.Delete)
}

func NewUserRoute(r *fiber.App, db *gorm.DB) {
	userService := service.UserService{
		Db: db,
	}

	userController := controller.UserController{
		Service: userService,
	}

	router := UserRouter{controller: userController}
	router.Load(r)
}
