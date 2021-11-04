package routes

import (
	"lab11/types/note"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
)

func PrepareRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("templates/*.tmpl")

	notes := note.CreateMockArray(10)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.tmpl", gin.H{
			"test": "test",
		})
	})

	router.GET("/notes", func(c *gin.Context) {
		tagsString := c.Query("tags")
		tagsArray := strings.Split(tagsString, ",")
		tagsQuery := mapset.NewSet()
		for _, item := range tagsArray {
			tagsQuery.Add(item)
		}
		c.HTML(http.StatusOK, "notes.tmpl", gin.H{
			"notes": notes,
		})
	})
	router.GET("/notes/:id", func(c *gin.Context) {
		itemId := c.Param("id")
		itemInt, err := strconv.ParseInt(itemId, 10, 0)
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "note.tmpl", gin.H{
			"note": notes[int(itemInt)],
		})
	})

	router.DELETE("/notes/:id", func(c *gin.Context) {
		itemId := c.Param("id")
		itemInt, err := strconv.ParseInt(itemId, 10, 0)
		if err != nil {
			panic(err)
		}
		notes = append(notes[:itemInt], notes[itemInt+1:]...)
		location := url.URL{Path: "/notes"}
		c.Redirect(http.StatusFound, location.RequestURI())
	})
}
