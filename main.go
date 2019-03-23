package main
import (
	"food-delivery/database"

	_ "github.com/jinzhu/gorm/dialects/postgres"

)
func main() {

	////initialize the database
	database.InitDB()


	///finally close the connection when you are done
	defer database.DBCon.Close()


}

