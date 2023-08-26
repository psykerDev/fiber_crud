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
	//user end points
	app.Post("/user", routes.CreateUser)
	app.Get("/users", routes.GetUsers)
	app.Get("/user/:id", routes.GetUserById)
	app.Patch("/user/:id", routes.UpdateUser)
	app.Delete("/user/:id", routes.DeleteUser)
	// post end points
	app.Post("/post", routes.CreatePost)
	app.Get("/posts", routes.GetPosts)
	app.Get("/post/:id", routes.GetPostById)
	app.Delete("/post/:id", routes.DeletePost)
	app.Patch("/post/:id", routes.UpdatePost)
	// reply end points
	app.Post("/reply", routes.CreateReply)
	app.Get("/reply", routes.GetReplys)
	app.Get("/reply/:id", routes.GetReplyById)
	app.Patch("/reply/:id", routes.UpdateReply)
	app.Delete("/reply/:id", routes.DeleteReply)
}
func main() {
	app := fiber.New()
	SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
