package routes

import "github.com/gin-gonic/gin"


type Router struct {
	server *gin.Engine
}

func (r *Router) SetRoutes() {
	warehouse := r.server.Group("/warehouse")

	//instances
	warehouse.GET("/reportProducts",)
}