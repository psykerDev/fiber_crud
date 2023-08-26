package routes

import (
	"errors"

	"main.go/initializers"
	"main.go/models"
)

// find user utility function
func FindUser(id int, user *models.User) error {
	initializers.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil

}
func FindPosts(id int, post *models.Post) error {
	initializers.DB.Find(&post, "id = ?", id)
	if post.ID == 0 {
		return errors.New("post does not exist")
	}
	return nil
}
func FindReply(id int, reply *models.Reply) error {

	initializers.DB.Find(&reply, "id = ?", id)
	if reply.ID == 0 {
		return errors.New("reply does not exist")
	}
	return nil
}
