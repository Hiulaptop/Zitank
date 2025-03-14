package models

import (
	"github.com/jackc/pgtype"
)

type Posts struct {
	ID         uint             `db:"id" json:"id"`
	Title      string           `db:"title" json:"title"`
	Content    string           `db:"content" json:"content"`
	CreateDate pgtype.Timestamp `db:"createdate" json:"createdate"`
	EditDate   pgtype.Timestamp `db:"editdate" json:"editdate"`
	UserID     uint             `db:"userid" json:"userid"`
}

type PostRepository interface {
	GetPosts() ([]*Posts, error)
	GetPost(int) (*Posts, error)
	GetAllPostByUser(int) ([]*Posts, error)

	CreatePost(*Posts) (int, error)

	UpdatePost(*Posts) error

	DeletePost(int) error
}
