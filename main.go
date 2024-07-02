package main

import (
	"fmt"
	"log"
	// database "todolist/Database"
	route "todolist/Route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Hi")
	// database.DBconnect()

	app := fiber.New()
	corsConfig := cors.Config{
		AllowOrigins:     "http://localhost:3000,http://172.20.10.11:8081,http://192.168.56.1:8081", // Replace with your frontend URL
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}
	app.Use(cors.New(corsConfig))
	route.SetUP(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal("server couldn't started " + err.Error())
	}
}
