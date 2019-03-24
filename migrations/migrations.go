package migrations

import (
	"food-delivery/database"
	"food-delivery/models"
)

func Migrate() {

	database.DBCon.AutoMigrate(models.User{}, models.Food{}, models.Address{}, models.OrderItems{}, models.Orders{})

	database.DBCon.Model(&models.User{}).AddForeignKey("address_id", "address(id)", "CASCADE", "RESTRICT")
	database.DBCon.Model(&models.Orders{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	database.DBCon.Model(&models.OrderItems{}).AddForeignKey("order_id", "orders(id)", "CASCADE", "RESTRICT")
	database.DBCon.Model(&models.OrderItems{}).AddForeignKey("food_id", "foods(id)", "CASCADE", "RESTRICT")


}
