package handlers

import (
	ServProd "clase2_web/internal/Productos"
	"clase2_web/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct{
	MetodoService ServProd.MetodosProductosService
}

//CONSTRUCTOR
func NuevoControlador(MetodoServ ServProd.MetodosProductosService) *Controller{
	return &Controller{MetodoService: MetodoServ}
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK,"pong")
}

func (con *Controller)GetAllProductos(c *gin.Context) {
	TodosProductos := con.MetodoService.GetAllService()
	c.JSON(200,TodosProductos)
}

func (con *Controller)GetByIdProducto(c *gin.Context) {
	Id, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil{
		c.JSON(http.StatusBadRequest,"error 1, al obtener el ID")
		return
	} else {
		Producto, err := con.MetodoService.GetByIdService(Id)
		if err != nil {
			c.JSON(http.StatusBadRequest,"error 2, al obtener la información de ese ID")
			return
		}
		c.JSON(200,Producto)
	}
}

func (con *Controller)GetMayorQueProducto(c *gin.Context) {
	min, err := strconv.ParseFloat(c.Query("priceGt"), 64)
	if err != nil{
		c.JSON(http.StatusBadRequest,err)
		return
	} else {
		Producto := con.MetodoService.GetMayorService(min)
		c.JSON(200,Producto)
	}
}

func (con *Controller)CrearNuevoProducto(c *gin.Context) {
	var request domain.Productos

	err1 := c.ShouldBind(&request)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, "error, Producto incorrecto")
		return
	}

	err2 := con.MetodoService.CrearNuevoProductoService(request)
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err2)
		return
	} else {
		c.JSON(201, "Producto guardado exitosamente")
	}

}