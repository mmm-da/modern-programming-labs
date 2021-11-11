package routes

import (
	"lab11/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PrepareRoutes(router *gin.Engine) {
	notes := make([]models.Note, 0)

	for i := 0; i < 10; i++ {
		notes = append(notes, models.CreateMockNote())
	}

	router.LoadHTMLGlob("templates/*.tmpl")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"notes": notes,
		})
	})
}
