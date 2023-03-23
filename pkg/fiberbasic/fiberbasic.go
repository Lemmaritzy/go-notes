package fiberbasic

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Mail        string `json:"mail"`
	PhoneNumber string `json:"phone_number"`
}

func FiberApp() {
	app := fiber.New()

	//For enabling cors
	app.Use(cors.New())

	routes(app)

	app.Listen(":8080")
}

func routes(c *fiber.App) {
	c.Get("/fiber-basics", mainHandler)
	c.Get("/send-data", sendData)
}

func mainHandler(c *fiber.Ctx) error {
	fmt.Println("hello web")
	return c.SendString("hello world ")

}

func sendData(c *fiber.Ctx) error {

	var data []User

	u1 := User{
		FirstName:   "mert",
		LastName:    "yilmaz",
		Mail:        "mert@test.com",
		PhoneNumber: "123 456 78 90",
	}

	data = append(data, u1)

	res, err := json.Marshal(data)

	if err != nil {
		return c.JSON(fiber.Map{
			"status": 500,
			"err":    err,
			"msg":    "an error occurred",
		})
	}

	return c.Send(res)
}
