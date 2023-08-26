package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/initializers"
	"main.go/models"
	routes "main.go/tools"
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

func GetReplys(c *fiber.Ctx) error {
	var post models.Post
	var postUser models.User
	var user models.User
	replys := []models.Reply{}

	responsReplys := []Reply{}

	initializers.DB.Find(&replys)

	for _, reply := range replys {
		if err := routes.FindPosts(int(reply.Post_ID), &post); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		if err := routes.FindUser(int(post.User_ID), &postUser); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		if err := routes.FindUser(int(reply.User_ID), &user); err != nil {
			return c.Status(400).JSON(err.Error())
		}
		responsPostUser := CreateResponsUser(postUser)
		responsUser := CreateResponsUser(user)
		responsPost := CreateResponsPost(post, responsPostUser)
		responsReply := CreateResponsReply(reply, responsPost, responsUser)
		responsReplys = append(responsReplys, responsReply)

	}
	return c.Status(200).JSON(responsReplys)

}
func GetReplyById(c *fiber.Ctx) error {
	var post models.Post
	var postUser models.User
	var user models.User
	var reply models.Reply
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("shit id")
	}
	if err := routes.FindReply(id, &reply); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindPosts(int(reply.Post_ID), &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(post.User_ID), &postUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(reply.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responsPostUser := CreateResponsUser(postUser)
	responsUser := CreateResponsUser(user)
	responsPost := CreateResponsPost(post, responsPostUser)
	responsReply := CreateResponsReply(reply, responsPost, responsUser)

	return c.Status(200).JSON(responsReply)

}
func UpdateReply(c *fiber.Ctx) error {
	var post models.Post
	var postUser models.User
	var user models.User
	var reply models.Reply
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("shit id")
	}
	if err := routes.FindReply(id, &reply); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindPosts(int(reply.Post_ID), &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(post.User_ID), &postUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(reply.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type updateReply struct {
		Comment string `json:"comment"`
	}

	var updateData updateReply

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	reply.Comment = updateData.Comment

	initializers.DB.Save(&reply)

	responsPostUser := CreateResponsUser(postUser)
	responsUser := CreateResponsUser(user)
	responsPost := CreateResponsPost(post, responsPostUser)
	responsReply := CreateResponsReply(reply, responsPost, responsUser)

	return c.Status(200).JSON(responsReply)

}
func DeleteReply(c *fiber.Ctx) error {
	var post models.Post
	var postUser models.User
	var user models.User
	var reply models.Reply
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(400).JSON("shit id")
	}
	if err := routes.FindReply(id, &reply); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindPosts(int(reply.Post_ID), &post); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(post.User_ID), &postUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := routes.FindUser(int(reply.User_ID), &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	initializers.DB.Delete(&reply, id)

	responsPostUser := CreateResponsUser(postUser)
	responsUser := CreateResponsUser(user)
	responsPost := CreateResponsPost(post, responsPostUser)
	responsReply := CreateResponsReply(reply, responsPost, responsUser)

	return c.Status(200).JSON(responsReply)
}
