package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
	database "todolist/Database"
	models "todolist/Models"
	util "todolist/Util"

	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func ValidateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	regex := regexp.MustCompile(emailRegex)

	return regex.MatchString(email)
}

func Register(ctx *fiber.Ctx) error {
	var data map[string]interface{}
	var userData models.User

	if err := ctx.BodyParser(&data); err != nil {
		log.Fatal("error to bind request body to struct " + err.Error())
	}

	if len(data["password"].(string)) <= 6 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "password must be greater than 6 character",
			"status":  400,
		})
	}
	if !ValidateEmail(strings.TrimSpace(data["email"].(string))) {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "Email is not valid",
			"status":  400,
		})
	}
	database.DB.Where("email = ?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id != 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "User already exist",
			"status":  400,
		})
	}

	user := models.User{
		Name:  data["name"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}
	user.SetPassword(data["password"].(string))
	database.DB.Create(&user)

	ctx.Status(200)
	return ctx.JSON(fiber.Map{
		"user":    user,
		"message": "Account created successfully",
		"status":  200,
	})
}

func Login(ctx *fiber.Ctx) error {
	var data = make(map[string]interface{})
	var userData models.User

	err := ctx.BodyParser(&data)
	if err != nil {
		fmt.Println("error occured during parsing of login data")
		log.Fatal(err)
	}

	database.DB.Where("email=?", data["email"].(string)).First(&userData)
	if userData.Id == 0 {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "User is not registered! please register first",
			"status":  400,
		})
	}
	err = userData.ComparePassword(data["password"].(string))
	if err != nil {
		ctx.Status(400)
		return ctx.JSON(fiber.Map{
			"message": "Password is incorrect",
			"status":  400,
		})
	}

	token, err := util.GenerateJWT(strconv.Itoa(userData.Id))
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return err
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Secure:   true,
	}

	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{
		"message": "You have successfully logged in",
		"user":    userData,
		"status":  200,
		"token":   token,
	})
}
