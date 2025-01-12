package controller

import (
	"Zitank/models"
)

type OrderController struct{}

func (OrderController) GetOrders(db *models.PostgresStore) ([]models.Orders, error) {
	var orders []models.Orders
	err := db.DB.Select(&orders, "SELECT * FROM orders")
	return orders, err
}

func (OrderController) GetOrder(db *models.PostgresStore, id int) (models.Orders, error) {
	var order models.Orders
	err := db.DB.Get(&order, "SELECT * FROM orders WHERE id=? LIMIT 1", id)
	return order, err
}

func (OrderController) CreateOrder(db *models.PostgresStore, order models.Orders) error {
	_, err := db.DB.Exec("INSERT INTO orders (user_id, room_id, check_in, check_out, total_price) VALUES (?, ?, ?, ?, ?)", order.UserID, order.RoomID, order.CheckInDate, order.CheckOutDate, order.TotalPrice)
	return err
}

func (OrderController) UpdateOrder(db *models.PostgresStore, order models.Orders) error {
	_, err := db.DB.Exec("UPDATE orders SET user_id=?, room_id=?, check_in=?, check_out=?, total_price=? WHERE id=?", order.UserID, order.RoomID, order.CheckInDate, order.CheckOutDate, order.TotalPrice, order.ID)
	return err
}

func (OrderController) DeleteOrder(db *models.PostgresStore, id int) error {
	_, err := db.DB.Exec("DELETE FROM orders WHERE id=?", id)
	return err
}

func (OrderController) GetAllOrderFromUser(db *models.PostgresStore, userID int) ([]models.Orders, error) {
	var orders []models.Orders
	err := db.DB.Select(&orders, "SELECT * FROM orders WHERE user_id=?", userID)
	return orders, err
}

func (OrderController) GetAllOrderFromRoom(db *models.PostgresStore, RoomID int) ([]models.Orders, error) {
	var orders []models.Orders
	err := db.DB.Select(&orders, "SELECT * FROM orders WHERE RoomID=?", RoomID)
	return orders, err
}
