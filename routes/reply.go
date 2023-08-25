package routes

import (
	"github.com/gofiber/fiber/v2"
	routes "main.go/Tooling"
	"main.go/initializers"
	"main.go/models"
)

type Reply struct {
	ID      uint   `json:"id"`
	Posts   Posts  `json:"post"`
	User    User   `json:"user"`
	Comment string `json:"comment"`
}

func CreateResponsReply(modelReply models.Reply, post Posts, user User) Reply {
	return Reply{ID: modelReply.ID, Posts: post, Comment: modelReply.Comment, User: user}
}
func CreateReply(c *fiber.Ctx) error {

	var posts models.Post
	var reply models.Reply
	var user models.User
	var postUser models.User
	if err := c.BodyParser(&reply); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindPosts(int(reply.Post_ID), &posts); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(reply.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(reply.Post_ID), &postUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	initializers.DB.Create(&reply)

	responsUser := CreateResponsUser(user)
	responsPostUser := CreateResponsUser(postUser)
	responsPost := CreateResponsPost(posts, responsPostUser)

	responseReply := CreateResponsReply(reply, responsPost, responsUser)
	return c.Status(200).JSON(responseReply)
}
