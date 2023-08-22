package routes

import "github.com/gofiber/fiber/v2"

type posts struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Greeting(c *fiber.Ctx) error {
	return c.SendString("Hello my friend")

}
