package routes

import (
	"time"

	"github.com/firdisml/go-rest/database"
	"github.com/firdisml/go-rest/models"
	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"order_date"`
	Product   Product   `json:"product"`
	User      User      `json:"user"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	return c.Status(200).JSON("Order Created")
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}

	database.Database.Db.Preload("Product").Preload("User").Find(&orders)

	return c.Status(200).JSON(orders)
}
