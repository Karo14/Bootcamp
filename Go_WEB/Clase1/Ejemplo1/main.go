package main

import "github.com/gin-gonic/gin"

func main(){
	//crea un router con gin
	router := gin.Default()

	//CAPTURA LA SOLICITUD GET "hola mundo"
	router.GET("Hola Mundo", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "Hola Mundo!",
		})
	})

	//corremos nuestro servidor sobre el puerto 8080
	router.Run()
}