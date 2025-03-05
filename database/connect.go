package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(DbName string, DbUser string, DbTcp string) *gorm.DB {
	dbName := DbName
	dbUser := DbUser
	dbTcp := DbTcp

	dsn := dbUser + dbTcp + dbName + "?charset=utf8&parseTime=True"

	gormDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Errore nella connessione al database: %v", err)
		return nil
	}

	fmt.Println("Connessione al database effettuata")

	return gormDb
}
