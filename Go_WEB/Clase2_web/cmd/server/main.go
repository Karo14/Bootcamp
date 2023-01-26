package main

import (
	"github.com/gin-gonic/gin"
	router "clase2_web/cmd/server/router"
)

func main() {
	servidor := gin.Default()
	router.Rutas(servidor)
	servidor.Run(":8080")
}