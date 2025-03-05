package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SalutaHandler(c *gin.Context) {
	nome := c.DefaultQuery("nome", "Guest")
	lavoro := c.DefaultQuery("lavoro", "Studente")

	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"nome":   nome,
		"lavoro": lavoro,
	})
}
