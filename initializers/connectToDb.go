package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/models"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "host=localhost user=postgres password= havlock dbname=go_blog  port=5432 sslmode=disable "
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
func SyncDb() {
	DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Reply{})
}