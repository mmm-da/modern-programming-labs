package main

import (
	"lab11/routes"

	"github.com/gin-gonic/gin"
)

// Задание
// CRUD для записок и фильтры по тегам

func main() {
	r := gin.Default()
	routes.PrepareRoutes(r)
	r.Run()
}
