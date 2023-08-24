package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"main.go/initializers"
	"main.go/routes"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDb()
}
func SetupRoutes(app *fiber.App) {
	app.Post("/", routes.CreateUser)
	app.Get("/allUsers", routes.GetUsers)
	app.Get("/getUser/:id", routes.GetUserById)
	app.Patch("/update/:id", routes.UpdateUser)
	app.Delete("/delUser/:id", routes.DeleteUser)
	app.Post("/createPost/:id", routes.CreatePost)
	app.Post("/reply/:id", routes.CreateReply)
}
func main() {
	app := fiber.New()
	SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
