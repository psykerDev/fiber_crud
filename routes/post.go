package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"main.go/initializers"
	"main.go/models"
)

type Posts struct {
	PostID  uint   `json:"post_id"`
	Title   string `json:"title"`
	Body    string `json:"body"`
	User_ID uint   `json:"user_id"`
}

func CreateresponsPost(postModel models.Post) Posts {
	return Posts{PostID: postModel.PostID, Title: postModel.Title, Body: postModel.Body, User_ID: postModel.User_ID}

}
func FindPosts(id int, post *models.Post) error {
	initializers.DB.Find(&post, "id = ?", id)
	if post.PostID == 0 {
		return errors.New("post does not exist")
	}
	return nil
}

// create post
func CreatePost(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	var posts models.Post

	if err != nil {
		return c.Status(400).JSON("invalid id ")
	}
	if err := findUser(id, &user); err != nil {
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
