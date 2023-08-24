package routes

import (
	"github.com/gofiber/fiber/v2"
	routes "main.go/Tooling"
	"main.go/initializers"
	"main.go/models"
)

type Posts struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	User_ID uint   `json:"user_id"`
}

func CreateresponsPost(postModel models.Post) Posts {
	return Posts{ID: postModel.ID, Title: postModel.Title, Body: postModel.Body, User_ID: postModel.User_ID}

}

// create post
func CreatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	var posts models.Post

	if err != nil {
		return c.Status(400).JSON("invalid id ")
	}
	if err := routes.FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := c.BodyParser(&posts); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Create(&posts)

	posts.User_ID = uint(id)
	initializers.DB.Save(&posts)
	responsPost := CreateresponsPost(posts)
	return c.Status(200).JSON(responsPost)

}

// get all post

// get post by id

// delete post

// update post
