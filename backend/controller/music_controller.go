package controller

import (
	"Zitank/models"
)

type MusicController struct{}

func (MusicController) GetMusicsByAlbumID(db *models.PostgresStore, id int) ([]models.Musics, error) {
	var music []models.Musics
	err := db.DB.Select(&music, "SELECT * FROM musics WHERE AlbumID=?", id)
	return music, err
}
func (MusicController) GetMusics(db *models.PostgresStore) ([]models.Musics, error) {
	var music []models.Musics
	err := db.DB.Select(&music, "SELECT * FROM musics")
	return music, err
}

func (MusicController) GetMusicsByID(db *models.PostgresStore, id int) (models.Musics, error) {
	var music models.Musics
	err := db.DB.Get(&music, "SELECT * FROM musics WHERE ID=? LIMIT 1", id)
	return music, err
}

func (MusicController) GetMusicInfoByMusicID(db *models.PostgresStore, id int) ([]models.MusicInfo, error) {
	var musicinfo []models.MusicInfo
	err := db.DB.Select(&musicinfo, "SELECT * FROM musicinfo WHERE MusicID=?", id)
	return musicinfo, err
}

func (MusicController) CreateMusic(db *models.PostgresStore, music models.Musics) error {
	_, err := db.DB.Exec("INSERT INTO musics (Name, Type, Link, AlbumID) VALUES (?, ?, ?, ?)", music.Name, music.Type, music.Link, music.AlbumID)
	return err
}

func (MusicController) CreateMusicInfo(db *models.PostgresStore, musicinfo models.MusicInfo) error {
	_, err := db.DB.Exec("INSERT INTO musicinfo (ArtistName, Role, Type, MusicID) VALUES (?, ?, ?, ?)", musicinfo.ArtistName, musicinfo.Role, musicinfo.Type, musicinfo.MusicID)
	return err
}

func (MusicController) UpdateMusic(db *models.PostgresStore, music models.Musics) error {
	_, err := db.DB.Exec("UPDATE musics SET Name=?, Type=?, Link=?, AlbumID=? WHERE id=?", music.Name, music.Type, music.Link, music.AlbumID, music.ID)
	return err
}

func (MusicController) UpdateMusicInfo(db *models.PostgresStore, musicinfo models.MusicInfo) error {
	_, err := db.DB.Exec("UPDATE musicinfo SET ArtistName=?, Role=?, Type=?, MusicID=? WHERE id=?", musicinfo.ArtistName, musicinfo.Role, musicinfo.Type, musicinfo.MusicID, musicinfo.ID)
	return err
}

func (MusicController) DeleteMusicInfo(db *models.PostgresStore, id int) error {
	_, err := db.DB.Exec("DELETE FROM musicinfo WHERE id=?", id)
	return err
}

func (MusicController) DeleteMusic(db *models.PostgresStore, id int) error {
	_, err := db.DB.Exec("DELETE FROM musics WHERE ID=?", id)
	return err
}
