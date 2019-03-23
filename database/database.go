package database

import "github.com/jinzhu/gorm"

var (

	DBCon *gorm.DB


)

func InitDB() {
	var err error

	//connect to postgres database
	DBCon, err = gorm.Open("postgres", "host=localhost port=myport user=gorm dbname=delivery password=mypassword")

	//where myhost is port is the port postgres is running on
	//user is your postgres use name
	//password is your postgres password
	if err != nil {
		panic(err)

		panic("failed to connect database")
	}

}
