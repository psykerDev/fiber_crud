package routes

import (
	"github.com/gofiber/fiber/v2"
	"main.go/initializers"
	"main.go/models"
	routes "main.go/tools"
)

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email" `
	Name  string `json:"name"`
}

func CreateResponsUser(userModel models.User) User {
	return User{ID: userModel.ID, Email: userModel.Email, Name: userModel.Name}
}

// create user
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Create(&user)

	responsUser := CreateResponsUser(user)
	return c.Status(200).JSON(responsUser)

}

// get all users
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	initializers.DB.Find(&users)

	responsUsers := []User{}

	for _, user := range users {
		responsUser := CreateResponsUser(user)
		responsUsers = append(responsUsers, responsUser)
	}
	return c.Status(200).JSON(responsUsers)

}

// get user by id
func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("id is trash")
	}
	if err := routes.FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	responsUser := CreateResponsUser(user)
	return c.Status(200).JSON(responsUser)

}

// update user
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("id is trash")
	}

	if err := routes.FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type updateUser struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	var updateData updateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	user.Email = updateData.Email
	user.Name = updateData.Name

	initializers.DB.Save(&user)
	responsUser := CreateResponsUser(user)

	return c.Status(200).JSON(responsUser)

}

// delete user

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("id is trash")
	}

	if err := routes.FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Delete(&user, id)
	responsUser := CreateResponsUser(user)
	return c.Status(200).JSON(responsUser)

}
