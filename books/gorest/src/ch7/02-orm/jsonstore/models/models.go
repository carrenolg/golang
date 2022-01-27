package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Data      string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
	UserRefer uint
}

type User struct {
	gorm.Model
	Orders []Order `gorm:"foreignKey:UserRefer"`
	Data   string  `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

// GORM creates tables with plural names. Use this to suppress it.
func (User) TableName() string {
	return "user"
}

func (Order) TableName() string {
	return "order"
}

func InitDB() (*gorm.DB, error) {
	var err error
	//dsn := "postgres://admin:1q2w3e4r@172.19.0.2/storedb?sslmode=disable"
	dsn := "host=172.19.0.2 user=admin password=1q2w3e4r dbname=storedb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	} else {
		// The below AutoMigrate is equivalent to this
		// Migrate the schema
		db.AutoMigrate(&User{}, &Order{})
		return db, nil
	}
}
