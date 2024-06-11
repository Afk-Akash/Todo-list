package route

import (
	controller "todolist/Controller"
	todocontroller "todolist/TodoController"

	"github.com/gofiber/fiber/v2"
)

func SetUP(app *fiber.App) {
	app.Get("/api/hello", todocontroller.Hello)
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Get("api/getTodo", todocontroller.GetTODO)
	app.Get("api/mob-getTodo", todocontroller.GetTODOMobile)

	app.Post("api/addTodo", todocontroller.AddTodo)
	app.Post("api/mob-addTodo", todocontroller.AddTodoMobile)

	app.Put("api/update", todocontroller.UpdateTodo)
	app.Put("api/mob-update", todocontroller.UpdateTodoMobile)

	app.Delete("api/delete", todocontroller.DeleteTodo)
	app.Delete("api/mob-delete", todocontroller.DeleteTodoMobile)

}
