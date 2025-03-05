package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTodosOfUserHandler(context *gin.Context) {
	dbInterface, _ := context.Get("DbKey")

	if dbInterface == nil {
		context.JSON(500, "Database non disponibile")
		return
	}

	db, ok := dbInterface.(*gorm.DB)

	if !ok {
		context.JSON(500, "Errore di tipo nella conversione del database")
		return
	}

	var todos []Todo
	username := context.MustGet("Username")

	risultato := db.Where("owner = ?", username).Find(&todos)

	if risultato.Error != nil {
		log.Fatalf("Errore nel fetch delle todo: %v", risultato.Error)
		context.JSON(500, "Errore nel fetch delle todo")
		return
	}

	fmt.Println("Fetch delle todo effettuato con successo: ", todos)
	context.JSON(200, todos)
}
