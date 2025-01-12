package models

import (
	"github.com/jackc/pgtype"
)

type Musics struct {
	ID         uint             `db:"ID"`
	Name       pgtype.Varchar   `db:"Name"`
	Type       pgtype.Varchar   `db:"Type"`
	Link       pgtype.Varchar   `db:"Link"`
	CreateDate pgtype.Timestamp `db:"CreateDate"`
	EditDate   pgtype.Timestamp `db:"EditDate"`
	AlbumID    uint             `db:"AlbumID"`
}

type MusicInfo struct {
	ID         uint             `db:"ID"`
	ArtistName pgtype.Varchar   `db:"ArtistName"`
	Role       pgtype.Varchar   `db:"Role"`
	Type       pgtype.Varchar   `db:"Type"`
	CreateDate pgtype.Timestamp `db:"CreateDate"`
	EditDate   pgtype.Timestamp `db:"EditDate"`
	MusicID    uint             `db:"MusicID"`
}
