package todocontroller

import (
	"fmt"
	"strconv"
	database "todolist/Database"
	models "todolist/Models"
	util "todolist/Util"

	"github.com/gofiber/fiber/v2"
)

func Hello(ctx *fiber.Ctx) error {
	ctx.Status(400)
	return ctx.JSON(fiber.Map{
		"message": "Hi...How are you this is testing api",
		"status":  200,
	})
}

func AddTodo(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	fmt.Println(ctx.Cookies("jwtToken"))
	id, _ := util.ParseJWT(cookie)
	var todo models.Todo

	if err := ctx.BodyParser(&todo); err != nil {
		fmt.Println("error occured during parsing add todo request ", err.Error())
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "internal server error- please try again later",
			"status":  400,
		})

	}

	todo.UserID, _ = strconv.Atoi(id)
	result := database.DB.Create(&todo)
	if result.Error != nil {
		ctx.Status(400)
		fmt.Println(result.Error)
		return ctx.JSON(fiber.Map{
			"message": "internal server error- please try again later",
			"status":  400,
		})
	}
	// result.val
	if result.RowsAffected == 0 {
		ctx.Status(400)
		fmt.Println("No rows were affected, couldn't add your todo")
		return ctx.JSON(fiber.Map{
			"message": "couldn't add your todo",
			"status":  400,
		})
	}

	ctx.Status(200)
	return ctx.JSON(fiber.Map{
		"message": "Added successfully",
		"status":  200,
		"id":      todo.Id,
	})
}

func GetTODO(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	fmt.Println(cookie)
	id, err := util.ParseJWT(cookie)
	if err != nil {
		fmt.Println("error occured during parsing jWT ", err.Error())
	}

	var todoList []models.Todo
	fmt.Println("User Id of Logged in User :-", id)
	database.DB.Where("user_id=?", id).Preload("User").Find(&todoList)
	return ctx.JSON(todoList)
}

func UpdateTodo(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	id, err := util.ParseJWT(cookie)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "Please login again...",
			"status":  400,
		})
	}
	var data = make(map[string]interface{})

	if err := ctx.BodyParser(&data); err != nil {
		fmt.Println("couldn't parse body of request in UpdateTODO ", err)
	}
	result := database.DB.Model(&models.Todo{}).Where("id = ?", data["id"]).Where("user_id=?", id).Updates(data)
	if result.Error != nil || result.RowsAffected == 0 {
		fmt.Println(result.Error)
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "updation failed",
			"status":  400,
		})
	}

	ctx.Status(200)
	return ctx.JSON(fiber.Map{
		"message": "updated successfully",
		"status":  200,
	})

}

func DeleteTodo(ctx *fiber.Ctx) error {
	cookie := ctx.Cookies("jwt")
	id, err := util.ParseJWT(cookie)
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "Please login again...",
			"status":  400,
		})
	}

	var data = make(map[string]interface{})

	if err := ctx.BodyParser(&data); err != nil {
		fmt.Println("couldn't parse body of request in UpdateTODO ", err)
	}

	todo := models.Todo{
		Id: int(data["id"].(float64)),
	}

	result := database.DB.Model(&models.Todo{}).Where("user_id=?", id).Delete(&todo)

	if result.RowsAffected == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "record not found",
			"status":  400,
		})
	}

	ctx.Status(200)
	return ctx.JSON(fiber.Map{
		"message": "deleted successfully",
		"status":  200,
	})
}
