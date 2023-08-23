package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"main.go/initializers"
	"main.go/models"
)

type User struct {
	ID    uint   `json:"id"`
	Email string `json:"email" `
	Name  string `json:"name"`
}

func CreateresponsUser(userModel models.User) User {
	return User{ID: userModel.ID, Email: userModel.Email, Name: userModel.Name}
}

// create user
func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Create(&user)

	responsUser := CreateresponsUser(user)
	return c.Status(200).JSON(responsUser)

}

// get all users
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	initializers.DB.Find(&users)

	responsUsers := []User{}

	for _, user := range users {
		responsUser := CreateresponsUser(user)
		responsUsers = append(responsUsers, responsUser)
	}
	return c.Status(200).JSON(responsUsers)

}

// find user utility function
func findUser(id int, user *models.User) error {
	initializers.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil

}

// get user by id
func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.User

	if err != nil {
		return c.Status(400).JSON("id is trash")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())

	}
	responsUser := CreateresponsUser(user)
	return c.Status(200).JSON(responsUser)

}

// update user
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("id is trash")
	}

	if err := findUser(id, &user); err != nil {
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
	responsUser := CreateresponsUser(user)

	return c.Status(200).JSON(responsUser)

}

// delete user

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("id is trash")
	}

	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	initializers.DB.Delete(&user, id)
	responsUser := CreateresponsUser(user)
	return c.Status(200).JSON(responsUser)

}
