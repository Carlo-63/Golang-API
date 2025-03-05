package handlers

import (
	"main/models"
	"main/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllUsers(db *gorm.DB) []User {
	var users []User
	db.Find(&users)
	return users
}

func RegisterHandler(context *gin.Context) {
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

	users := GetAllUsers(db)
	username := context.Query("username")
	password := context.Query("password")
	presente := false

	for _, user := range users {
		if user.Username == username {
			presente = true
			break
		}
	}

	if presente {
		context.JSON(400, "Username già presente")
		return
	}

	secretKey := context.MustGet("ConfigKey").(models.Config).SecretKey

	db.Create(&User{Username: username, Password: password})

	token, err := utils.GenerateToken(username, []byte(secretKey))

	if err != nil {
		context.JSON(500, "Errore nella generazione del token")
		return
	}
	context.JSON(202, gin.H{"token": token})
}
