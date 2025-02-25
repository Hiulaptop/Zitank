package models

import (
	"github.com/jackc/pgtype"
)

type Musics struct {
	ID         uint             `db:"id" json:"id"`
	Name       string           `db:"name" json:"name"`
	Type       string           `db:"type" json:"type"`
	Link       string           `db:"link" json:"link"`
	CreateDate pgtype.Timestamp `db:"createdate" json:"createdate"`
	EditDate   pgtype.Timestamp `db:"editdate" json:"editdate"`
	AlbumID    uint             `db:"albumid" json:"albumid"`
}

type MusicInfo struct {
	ID         uint             `db:"id" json:"id"`
	ArtistName string           `db:"artistname"  json:"artistname"`
	Role       string           `db:"role"  json:"role"`
	Type       string           `db:"type"  json:"type"`
	CreateDate pgtype.Timestamp `db:"createdate"  json:"createdate"`
	EditDate   pgtype.Timestamp `db:"editdate"  json:"editdate"`
	MusicID    uint             `db:"musicid"  json:"musicid"`
}

type MusicRepository interface {
	GetMusicsByAlbumID(int) ([]*Musics, error)
	GetMusics() ([]*Musics, error)
	GetMusicsByID(int) (*Musics, error)
	GetMusicInfoByMusicID(int) ([]*MusicInfo, error)

	CreateMusic(*Musics) error
	CreateMusicInfo(*MusicInfo) error

	UpdateMusic(*Musics) error
	UpdateMusicInfo(*MusicInfo) error

	DeleteMusic(int) error
	DeleteMusicInfo(int) error
}
