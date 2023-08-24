package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/initializers"
	"main.go/routes"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDb()

}
func main() {
	app := fiber.New()

	app.Post("/", routes.CreateUser)
	app.Get("/allUsers", routes.GetUsers)
	app.Get("/getUser/:id", routes.GetUserById)
	app.Patch("/update/:id", routes.UpdateUser)
	app.Delete("/delUser/:id", routes.DeleteUser)
	app.Post("/createPost/:id", routes.CreatePost)
	app.Post("/reply/:id", routes.CreateReply)

	app.Listen(":8080")
}
