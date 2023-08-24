package routes

import (
	"github.com/gofiber/fiber/v2"
	routes "main.go/Tooling"
	"main.go/initializers"
	"main.go/models"
)

type Reply struct {
	ID      uint   `json:"id"`
	Post_ID uint   `json:"post_id"`
	Comment string `json:"comment"`
}

func CreateResponsReply(modelReply models.Reply) Reply {
	return Reply{ID: modelReply.ID, Post_ID: modelReply.Post_ID, Comment: modelReply.Comment}
}
func CreateReply(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var posts models.Post
	var reply models.Reply

	if err != nil {
		return c.Status(400).JSON("shit id ")
	}
	if err := routes.FindPosts(id, &posts); err != nil {

		return c.Status(400).JSON(err.Error())
	}

	if err := c.BodyParser(&reply); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Create(&reply)

	reply.Post_ID = uint(id)
	initializers.DB.Save(&reply)

	responseReply := CreateResponsReply(reply)
	return c.Status(200).JSON(responseReply)
}
