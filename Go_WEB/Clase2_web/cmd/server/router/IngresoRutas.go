package router

import (
	"github.com/gin-gonic/gin"
	"clase2_web/internal/Productos"
	"clase2_web/cmd/server/handlers"
)

func Rutas(servidor *gin.Engine){

	ProductosRouter := servidor.Group("/Productos")

	Repo,_ := Productos.NuevoRepository()
	Serv := Productos.NuevoServicio(Repo)
	Contr := handlers.NuevoControlador(Serv)

	servidor.GET("/ping", handlers.Ping)
	ProductosRouter.GET("/", Contr.GetAllProductos)
	ProductosRouter.GET("/:id", Contr.GetByIdProducto)
	ProductosRouter.GET("/search", Contr.GetMayorQueProducto)
	ProductosRouter.POST("/", Contr.CrearNuevoProducto)
}