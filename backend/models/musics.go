package models

import (
	"github.com/jackc/pgtype"
)

type Musics struct {
	ID         uint             `db:"ID" json:"ID"`
	Name       pgtype.Varchar   `db:"Name" json:"Name"`
	Type       pgtype.Varchar   `db:"Type" json:"Type"`
	Link       pgtype.Varchar   `db:"Link" json:"Link"`
	CreateDate pgtype.Timestamp `db:"CreateDate" json:"CreateDate"`
	EditDate   pgtype.Timestamp `db:"EditDate" json:"EditDate"`
	AlbumID    uint             `db:"AlbumID" json:"AlbumID"`
}

type MusicInfo struct {
	ID         uint             `db:"ID" json:"ID"`
	ArtistName pgtype.Varchar   `db:"ArtistName"  json:"ArtistName"`
	Role       pgtype.Varchar   `db:"Role"  json:"Role"`
	Type       pgtype.Varchar   `db:"Type"  json:"Type"`
	CreateDate pgtype.Timestamp `db:"CreateDate"  json:"CreateDate"`
	EditDate   pgtype.Timestamp `db:"EditDate"  json:"EditDate"`
	MusicID    uint             `db:"MusicID"  json:"MusicID"`
}
