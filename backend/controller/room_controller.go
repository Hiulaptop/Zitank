package controller

import (
	"Zitank/models"
)

type RoomController struct{}

func (RoomController) GetRooms(db *models.PostgresStore) ([]models.Rooms, error) {
	var rooms []models.Rooms
	err := db.DB.Select(&rooms, "SELECT * FROM rooms")
	return rooms, err
}

func (RoomController) GetRoom(db *models.PostgresStore, id int) (models.Rooms, error) {
	var room models.Rooms
	err := db.DB.Get(&room, "SELECT * FROM rooms WHERE id=? LIMIT 1", id)
	return room, err
}

func (RoomController) CreateRoom(db *models.PostgresStore, room models.Rooms) error {
	_, err := db.DB.Exec("INSERT INTO rooms (name, address, description, price) VALUES (?, ?, ?, ?)", room.Name, room.Address, room.Description, room.Price)
	return err
}

func (RoomController) UpdateRoom(db *models.PostgresStore, room models.Rooms) error {
	_, err := db.DB.Exec("UPDATE rooms SET name=?, address=?, description=?, price=? WHERE id=?", room.Name, room.Address, room.Description, room.Price, room.ID)
	return err
}

func (RoomController) DeleteRoom(db *models.PostgresStore, id int) error {
	_, err := db.DB.Exec("DELETE FROM rooms, WHERE id=?", id)
	return err
}
