package main

import "github.com/gin-gonic/gin"

type Persona struct {
	Nombre string
	Apellido string
}

func main(){
	router:= gin.Default()

	router.POST("/saludo",func (c *gin.Context)  {
		body := Persona{}

		if err := c.ShouldBind(&body); err != nil {
			c.Error(err)
			c.AbortWithStatus(400)
		}

		c.String(200, "Hola "+body.Nombre+" "+body.Apellido)
	})

	router.Run(":8081")
}