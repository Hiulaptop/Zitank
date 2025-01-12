package controller

import (
	"Zitank/models"
)

type PostController struct{}

func (PostController) GetPosts(db *models.PostgresStore) ([]models.Posts, error) {
	var posts []models.Posts
	err := db.DB.Select(&posts, "SELECT * FROM posts")
	return posts, err
}

func (PostController) GetPost(db *models.PostgresStore, id int) (models.Posts, error) {
	var post models.Posts
	err := db.DB.Select(&post, "SELECT * FROM posts WHERE id=? LIMIT 1", id)
	return post, err
}

func (PostController) GetAllPostByUser(db *models.PostgresStore, uid int) ([]models.Posts, error) {
	var posts []models.Posts
	err := db.DB.Select(&posts, "SELECT * FROM posts WHERE UserID=?", uid)
	return posts, err
}

func (PostController) CreatePost(db *models.PostgresStore, post models.Posts) error {
	_, err := db.DB.Exec("INSERT INTO posts (Title, Content, UserID) VALUES (?, ?, ?)", post.Title, post.Content, post.UserID)
	return err
}

func (PostController) UpdatePost(db *models.PostgresStore, post models.Posts) error {
	_, err := db.DB.Exec("UPDATE posts SET Title=?, Content=? WHERE id=?", post.Title, post.Content, post.ID)
	return err
}

func (PostController) DeletePost(db *models.PostgresStore, id int) error {
	_, err := db.DB.Exec("DELETE FROM posts WHERE id=?", id)
	return err
}
