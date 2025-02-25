package repositories

import (
	"Zitank/models"

	"github.com/jmoiron/sqlx"
)

type MusicRepo struct {
	DB *sqlx.DB
}

func NewMusicRepo(db *sqlx.DB) *MusicRepo {
	return &MusicRepo{
		DB: db,
	}
}

func (MR MusicRepo) GetMusicsByAlbumID(id int) ([]*models.Musics, error) {
	var music []*models.Musics
	err := MR.DB.Select(&music, `SELECT * FROM musics WHERE albumid=$1`, id)
	return music, err
}
func (MR MusicRepo) GetMusics() ([]*models.Musics, error) {
	var music []*models.Musics
	err := MR.DB.Select(&music, `SELECT * FROM musics`)
	return music, err
}

func (MR MusicRepo) GetMusicsByID(id int) (*models.Musics, error) {
	var music *models.Musics
	err := MR.DB.Get(&music, `SELECT * FROM musics WHERE id=$1 LIMIT 1`, id)
	return music, err
}

func (MR MusicRepo) GetMusicInfoByMusicID(id int) ([]*models.MusicInfo, error) {
	var musicinfo []*models.MusicInfo
	err := MR.DB.Select(&musicinfo, `SELECT * FROM musicinfo WHERE musicid=$1`, id)
	return musicinfo, err
}

func (MR MusicRepo) CreateMusic(music *models.Musics) error {
	_, err := MR.DB.Exec(`INSERT INTO musics (name, type, link, albumid) VALUES ($1, $2, $3, $4)`, music.Name, music.Type, music.Link, music.AlbumID)
	return err
}

func (MR MusicRepo) CreateMusicInfo(musicinfo *models.MusicInfo) error {
	_, err := MR.DB.Exec(`INSERT INTO musicinfo (artistname, role, type, musicid) VALUES ($1, $2, $3, $4)`, musicinfo.ArtistName, musicinfo.Role, musicinfo.Type, musicinfo.MusicID)
	return err
}

func (MR MusicRepo) UpdateMusic(music *models.Musics) error {
	_, err := MR.DB.Exec(`UPDATE musics SET name=$1, type=$2, link=$3, albumid=$4 WHERE id=$5`, music.Name, music.Type, music.Link, music.AlbumID, music.ID)
	return err
}

func (MR MusicRepo) UpdateMusicInfo(musicinfo *models.MusicInfo) error {
	_, err := MR.DB.Exec(`UPDATE musicinfo SET artistname=$1, role=$2, type=$3, musicid=$4 WHERE id=$5`, musicinfo.ArtistName, musicinfo.Role, musicinfo.Type, musicinfo.MusicID, musicinfo.ID)
	return err
}

func (MR MusicRepo) DeleteMusicInfo(id int) error {
	_, err := MR.DB.Exec(`DELETE FROM musicinfo WHERE id=$1`, id)
	return err
}

func (MR MusicRepo) DeleteMusic(id int) error {
	_, err := MR.DB.Exec(`DELETE FROM musics WHERE id=$1`, id)
	return err
}
