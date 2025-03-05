package handlers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddTodoHandler(context *gin.Context) {
	dbInterface, _ := context.Get("DbKey")

	if dbInterface == nil {
		context.String(500, "Database non disponibile")
		return
	}

	db, ok := dbInterface.(*gorm.DB)

	if !ok {
		context.String(500, "Errore di tipo nella conversione del database")
		return
	}

	name := context.DefaultQuery("name", "Default")
	username := context.MustGet("Username")

	todo := Todo{
		Name:  name,
		Owner: username.(string),
	}

	risultato := db.Create(&todo)

	if risultato.Error != nil {
		log.Fatalf("Errore durante la creazione della todo: %v", risultato.Error)
		context.String(404, "Errore nella creazione della todo")
		return
	}
	fmt.Println("Nuova todo creata")

	context.String(200, "Todo creata")
}
