package handlers

import (
	productsServices "Clase3_web/internal/products"
	models "Clase3_web/models"
	//"errors"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

var Prods productsServices.Products

func CreateProduct(ctx *gin.Context) {

	var req models.Request
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	if product := Prods.GetByCodeValue(req.Code_value); product != nil {
		ctx.JSON(400, "El producto ya existe")
		return
	}

	if _, err := time.Parse("02/01/2006", req.Expiration); err != nil {
		ctx.JSON(400, "Formato de fecha inválido")
		return
	}

	Prods.Create(req)

	ctx.JSON(201, "Producto guardado exitosamente")
}

func GetAllHandler(ctx *gin.Context) {
	products := Prods.GetAll()
	ctx.JSON(200, products)
}

func GetPriceGreaterThanHandler(ctx *gin.Context) {
	min, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	products := Prods.GetGreaterThan(min)

	ctx.JSON(200, products)
}

func GetOneHandler(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	product := Prods.GetById(id)
	if product == nil {
		ctx.JSON(400, "Producto no encontrado")
		return
	}
	ctx.JSON(200, product)
}
