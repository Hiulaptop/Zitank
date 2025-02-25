package repositories

import (
	"Zitank/models"

	"github.com/jmoiron/sqlx"
)

type PostRepo struct {
	DB *sqlx.DB
}

func NewPostRepo(PR *sqlx.DB) *PostRepo {
	return &PostRepo{
		DB: PR,
	}
}

func (PR PostRepo) GetPosts() ([]*models.Posts, error) {
	var posts []*models.Posts
	err := PR.DB.Select(&posts, `SELECT * FROM posts`)
	return posts, err
}

func (PR PostRepo) GetPost(id int) (*models.Posts, error) {
	var post *models.Posts
	err := PR.DB.Select(&post, `SELECT * FROM posts WHERE id=$1 LIMIT 1`, id)
	return post, err
}

func (PR PostRepo) GetAllPostByUser(uid int) ([]*models.Posts, error) {
	var posts []*models.Posts
	err := PR.DB.Select(&posts, `SELECT * FROM posts WHERE userid=$1`, uid)
	return posts, err
}

func (PR PostRepo) CreatePost(post *models.Posts) error {
	_, err := PR.DB.Exec(`INSERT INTO posts (title, content, userid) VALUES ($1, $2, $3)`, post.Title, post.Content, post.UserID)
	return err
}

func (PR PostRepo) UpdatePost(post *models.Posts) error {
	_, err := PR.DB.Exec(`UPDATE posts SET title=$1, content=$2 WHERE id=$3`, post.Title, post.Content, post.ID)
	return err
}

func (PR PostRepo) DeletePost(id int) error {
	_, err := PR.DB.Exec(`DELETE FROM posts WHERE id=$1`, id)
	return err
}
