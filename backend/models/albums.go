package models

import (
	"github.com/jackc/pgtype"
)

type Albums struct {
	ID          uint             `db:"ID" json:"ID"`
	Name        string           `db:"Name" json:"Name"`
	AuthorName  string           `db:"AuthorName" json:"AuthorName"`
	ReleaseDate pgtype.Timestamp `db:"ReleaseDate" json:"ReleaseDate"`
	Type        string           `db:"Type" json:"Type"`
	Description pgtype.Text      `db:"Description" json:"Description"`
	Link        string           `db:"Link" json:"Link"`
	CreateDate  pgtype.Timestamp `db:"CreateDate" json:"CreateDate"`
	EditDate    pgtype.Timestamp `db:"EditDate" json:"EditDate"`
	UserID      uint             `db:"UserID" json:"UserID"`
}

type AlbumRepository interface {
	GetAlbums() ([]*Albums, error)
	GetAlbumsByID(int) (*Albums, error)
	GetAlbumsByName(string) ([]*Albums, error)
	GetAlbumsByAuthorName(string) ([]*Albums, error)

	CreateAlbum(*Albums) error

	UpdateAlbum(*Albums) error

	DeleteAlbum(int) error
}
