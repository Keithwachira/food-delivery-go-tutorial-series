package main
import (
	"food-delivery/database"
	"food-delivery/migrations"

	_ "github.com/jinzhu/gorm/dialects/postgres"

)
func main() {

	////initialize the database
	database.InitDB()

	migrations.Migrate()///add this line to main.go to initialize the migration


	///finally close the connection when you are done
	defer database.DBCon.Close()


}

