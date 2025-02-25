package repositories

import (
	"Zitank/models"

	"github.com/jmoiron/sqlx"
)

type AlbumRepo struct {
	DB *sqlx.DB
}

func NewAlbumRepo(db *sqlx.DB) *AlbumRepo {
	return &AlbumRepo{
		DB: db,
	}
}

func (AR *AlbumRepo) GetAlbums() ([]*models.Albums, error) {
	var albums []*models.Albums
	err := AR.DB.Select(&albums, `SELECT * FROM albums`)
	return albums, err
}

func (AR *AlbumRepo) GetAlbumsByID(id int) (*models.Albums, error) {
	var album *models.Albums
	err := AR.DB.Get(&album, `SELECT * FROM albums WHERE ID=$1 LIMIT 1`, id)
	return album, err
}

func (AR *AlbumRepo) GetAlbumsByName(name string) ([]*models.Albums, error) {
	var album []*models.Albums
	err := AR.DB.Select(&album, `SELECT * FROM albums WHERE Name=$1`, name)
	return album, err
}

func (AR *AlbumRepo) GetAlbumsByAuthorName(authorname string) ([]*models.Albums, error) {
	var album []*models.Albums
	err := AR.DB.Select(&album, `SELECT * FROM albums WHERE AuthorName=$1`, authorname)
	return album, err
}

func (AR *AlbumRepo) CreateAlbum(album *models.Albums) error {
	_, err := AR.DB.Exec(`INSERT INTO albums (Name, AuthorName, ReleaseDate, Type, Description, Link, UserID) VALUES ($1, $2, $3, $4, $5, $6, $7)`, album.Name, album.AuthorName, album.ReleaseDate, album.Type, album.Description, album.Link, album.UserID)
	return err
}

func (AR *AlbumRepo) UpdateAlbum(album *models.Albums) error {
	_, err := AR.DB.Exec(`UPDATE albums SET Name=$1, AuthorName=$2, ReleaseDate=$3, Type=$4, Description=$5, Link=$6 WHERE id=$7`, album.Name, album.AuthorName, album.ReleaseDate, album.Type, album.Description, album.Link, album.ID)
	return err
}

func (AR *AlbumRepo) DeleteAlbum(id int) error {
	_, err := AR.DB.Exec(`DELETE * FROM albums WHERE ID=$1`, id)
	return err
}
