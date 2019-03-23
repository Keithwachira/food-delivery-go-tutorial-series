package models

import "time"

type Food struct {
	//food will contains this fields....
	//we have also defined they type
	ID   int `json:"id,primary_key"` //this will be the primary key field
	Name int `gorm:"size:255"`       ///the name field will only allow a maximum of 255 characters

	OrderItems []OrderItems
	//this defines the relationship between orderitems and food
	//its says that one food can belong to many orderitems and foodId is the foreign_key
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	//this represents the customer
	ID int `json:"id,primary_key"` //this will be the primary key field

	FirstName  string `gorm:"size:255"`
	SecondName string `gorm:"size:255"`

	Orders []Orders
	//this defines the relationship between orders and users
	//its says that one user can have to many orders and UserID is the foreign_key in the order table
	AddressID int
	///address_id is a foreign_key from the address table....
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Orders struct {
	ID int `json:"id,primary_key"` //this will be the primary key field

	OrderItems []OrderItems
	//this defines the relationship between orders and orderitems
	//its says that one order can have  many orderitems and OrderId is the foreign_key in the orderitems table

	UserID    int       `json:"user_id"` ///this is a foreign_key from the user_table
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderItems struct {
	ID        int       `json:"id,primary_key"` //this will be the primary key field
	OrderId   int                               //this is foreign_key from the  order table
	FoodID    int                               //this is foreign_key from the  food table
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Address struct {
	ID   int `json:"id,primary_key"` //this will be the primary key field
	User []User
	//this defines the relationship between users  and Address
	//its says that one address can belong to many users and  is the foreign_key address_id
	AddressName string    `gorm:"size:255"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
