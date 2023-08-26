package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/initializers"
	"main.go/models"
	routes "main.go/tools"
)

type Posts struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	User  User   `json:"user"`
}

func CreateResponsPost(postModel models.Post, user User) Posts {
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

	responsUser := CreateResponsUser(user)
	responsPost := CreateResponsPost(posts, responsUser)
	return c.Status(200).JSON(responsPost)

}

// get all post
func GetPosts(c *fiber.Ctx) error {
	posts := []models.Post{}
	var user models.User

	initializers.DB.Find(&posts)

	responsPosts := []Posts{}

	for _, post := range posts {

		if err := routes.FindUser(int(post.User_ID), &user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responsUser := CreateResponsUser(user)
		responsPost := CreateResponsPost(post, responsUser)
		responsPosts = append(responsPosts, responsPost)
	}
	return c.Status(200).JSON(responsPosts)
}

// get post by id
func GetPostById(c *fiber.Ctx) error {
	var post models.Post
	var user models.User
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("invalid id")
	}

	if err := routes.FindPosts(id, &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(post.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responsUser := CreateResponsUser(user)
	responsPost := CreateResponsPost(post, responsUser)

	return c.Status(200).JSON(responsPost)

}

// delete post
func DeletePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var post models.Post
	var user models.User

	if err != nil {
		return c.Status(400).JSON("invalid id")
	}
	if err := routes.FindPosts(id, &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(post.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Delete(&post, id)
	responsUser := CreateResponsUser(user)
	responsPost := CreateResponsPost(post, responsUser)
	return c.Status(200).JSON(responsPost)
}

// update post
func UpdatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var post models.Post
	var user models.User
	if err != nil {
		return c.Status(400).JSON("invalid id ")
	}

	if err := routes.FindPosts(id, &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(post.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updatePost struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	var updateData updatePost

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	post.Title = updateData.Title
	post.Body = updateData.Body

	initializers.DB.Save(&post)
	responsUser := CreateResponsUser(user)
	responsPost := CreateResponsPost(post, responsUser)
	return c.Status(200).JSON(responsPost)
}
