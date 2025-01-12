package controller

import (
	"Zitank/models"
)

type AlbumController struct{}

func (AlbumController) GetAlbums(db *models.PostgresStore) ([]models.Albums, error) {
	var albums []models.Albums
	err := db.DB.Select(&albums, "SELECT * FROM albums")
	return albums, err
}

func (AlbumController) GetAlbumsByID(db *models.PostgresStore, id int) (models.Albums, error) {
	var album models.Albums
	err := db.DB.Get(&album, "SELECT * FROM albums WHERE ID=? LIMIT 1", id)
	return album, err
}

func (AlbumController) GetAlbumsByName(db *models.PostgresStore, name string) ([]models.Albums, error) {
	var album []models.Albums
	err := db.DB.Select(&album, "SELECT * FROM albums WHERE Name=?", name)
	return album, err
}

func (AlbumController) GetAlbumsByAuthorName(db *models.PostgresStore, authorname string) ([]models.Albums, error) {
	var album []models.Albums
	err := db.DB.Select(&album, "SELECT * FROM albums WHERE AuthorName=?", authorname)
	return album, err
}

func (AlbumController) CreateAlbum(db *models.PostgresStore, album models.Albums) error {
	_, err := db.DB.Exec("INSERT INTO albums (Name, AuthorName, ReleaseDate, Type, Description, Link, UserID) VALUES (?, ?, ?, ?, ?, ?, ?)", album.Name, album.AuthorName, album.ReleaseDate, album.Type, album.Description, album.Link, album.UserID)
	return err
}

func (AlbumController) UpdateAlbum(db *models.PostgresStore, album *models.Albums) error {
	_, err := db.DB.Exec("UPDATE albums SET Name=?, AuthorName=?, ReleaseDate=?, Type=?, Description=?, Link=? WHERE id=?", album.Name, album.AuthorName, album.ReleaseDate, album.Type, album.Description, album.Link, album.ID)
	return err
}

func (AlbumController) DeleteAlbum(db *models.PostgresStore, id int) error {
	_, err := db.DB.Exec("DELETE * FROM albums WHERE ID=?", id)
	return err
}
