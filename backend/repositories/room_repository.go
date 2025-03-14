package repositories

import (
	"Zitank/models"

	"github.com/jmoiron/sqlx"
)

type RoomRepo struct {
	DB *sqlx.DB
}

func NewRoomRepo(db *sqlx.DB) *RoomRepo {
	return &RoomRepo{
		DB: db,
	}
}

func (RR RoomRepo) GetRooms() ([]*models.Rooms, error) {
	var rooms []*models.Rooms
	err := RR.DB.Select(&rooms, `SELECT * FROM rooms`)
	return rooms, err
}

func (RR RoomRepo) GetRoom(id int) (*models.Rooms, error) {
	var room *models.Rooms
	err := RR.DB.Get(&room, `SELECT * FROM rooms WHERE id=$1 LIMIT 1`, id)
	return room, err
}

func (RR RoomRepo) CreateRoom(room *models.RoomObject, userID int) (int, error) {
	var id int
	err := RR.DB.QueryRow(`INSERT INTO rooms (name, address, description, price, userid) VALUES ($1, $2, $3, $4, $5) returning id;`, room.Name, room.Address, room.Description, room.Price, userID).Scan(&id)
	return id, err
}

func (RR RoomRepo) UpdateRoom(room *models.Rooms) error {
	_, err := RR.DB.Exec(`UPDATE rooms SET name=$1, address=$2, description=$3, price=$4 WHERE id=$5`, room.Name, room.Address, room.Description, room.Price, room.ID)
	return err
}

func (RR RoomRepo) DeleteRoom(id int) error {
	_, err := RR.DB.Exec(`DELETE FROM rooms, WHERE id=$1`, id)
	return err
}
