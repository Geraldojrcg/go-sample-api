package route

import (
	"github.com/geraldojrcg/go-sample-api/src/controller"
	"github.com/geraldojrcg/go-sample-api/src/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodoRouter struct {
	controller controller.TodoController
}

func (t *TodoRouter) Load(r *fiber.App) {
	r.Get("/todos", t.controller.GetAll)
	r.Get("/todos/:id", t.controller.GetById)
	r.Post("/todos", t.controller.Create)
	r.Patch("/todos/:id", t.controller.Update)
	r.Delete("/todos/:id", t.controller.Delete)
}

func NewTodoRoute(r *fiber.App, db *gorm.DB) {
	userService := service.UserService{
		Db: db,
	}

	todoService := service.TodoService{
		Db:          db,
		UserService: userService,
	}

	todoController := controller.TodoController{
		Service: todoService,
	}

	router := TodoRouter{controller: todoController}
	router.Load(r)
}
