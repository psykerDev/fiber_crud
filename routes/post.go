package routes

import (
	"github.com/gofiber/fiber/v2"
	routes "main.go/Tooling"
	"main.go/initializers"
	"main.go/models"
)

type Posts struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	User  User   `json:"user"`
}

func CreateresponsPost(postModel models.Post, user User) Posts {
	return Posts{ID: postModel.ID, Title: postModel.Title, Body: postModel.Body, User: user}

}

// create post
func CreatePost(c *fiber.Ctx) error {
	var user models.User
	var posts models.Post
	if err := c.BodyParser(&posts); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := routes.FindUser(int(posts.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	initializers.DB.Create(&posts)

	responsUser := CreateresponsUser(user)
	responsPost := CreateresponsPost(posts, responsUser)
	return c.Status(200).JSON(responsPost)

}

// get all post

// get post by id

// delete post

// update post
