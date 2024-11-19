package mysql

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMySQLConnection(dsn string) *gorm.DB {
	dsnParsed := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", dsn)
	db, err := gorm.Open(mysql.Open(dsnParsed), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Conexión a MySQL establecida")

	return db
}
