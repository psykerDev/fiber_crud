package initializers

import "main.go/models"

func SyncDb() {
	DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Reply{})
}
