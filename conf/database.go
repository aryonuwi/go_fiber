package conf

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DatabaseInit() {

	var err error

	const POSTGRES = "host=localhost user=postgres password=aryo1291 dbname=go_fiber port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	dsn := POSTGRES
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Can't connected to database!!!")
	}
	fmt.Println("***Database successfully connected***")

}
