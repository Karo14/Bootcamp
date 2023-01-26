package Productos

import (
	"clase2_web/internal/domain"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

type ProductosAll struct{
	ArrayProductos []domain.Productos
}

//crear una interface con todos los metodos 
type MetodosRepository interface{
	GetAllRepository() []domain.Productos
	GetByIdRepository(id int) (domain.Productos,error)
	GetMayorRepository(min float64) []domain.Productos
	CrearNuevoProductoRepository(data domain.Productos) error
	GetByCodeValue(code string) bool
}

var TodosProd ProductosAll

//CONSTRUCTOR
func NuevoRepository() (*ProductosAll, error){
	var arr []domain.Productos

	//leer archivo
	datos, err:= os.ReadFile("/Users/angcarrillo/Documents/Clase2_web/products.json")
	if err != nil {
		return &ProductosAll{}, err
	}

	//se deserealiza el json al modelo
	err = json.Unmarshal(datos, &arr)
	if err != nil {
		return &ProductosAll{}, err
	}
	return &ProductosAll{ArrayProductos: arr},nil
}

//mostrar todo los productos
func (p *ProductosAll)GetAllRepository() []domain.Productos {
	//se devuelve el array con todos los productos
	return p.ArrayProductos
}

//buscar un productos por un id
func (p ProductosAll) GetByIdRepository(id int) (domain.Productos, error) {
	var idretorno domain.Productos

	for _, v := range p.ArrayProductos {
		if v.Id == id {
			idretorno = v
			break
		}
	}

	if idretorno.Id != 0 {
		return idretorno,nil
	}

	return idretorno, errors.New("error al obtener el producto por ID")
}

//buscar por valor mayor que
func (p ProductosAll) GetMayorRepository(min float64) []domain.Productos {
	var res []domain.Productos

	for _, v := range p.ArrayProductos {
		if v.Price > min {
			res = append(res, v)
		}
	}
	return res
}

//crear nuevo producto
func (p *ProductosAll) CrearNuevoProductoRepository(data domain.Productos) error {
	var valid bool
	var err error

	valid, err = validateEmptys(&data)
	if(!valid){
		return err
	}

	valid, err = validateExpiration(&data)
	if(!valid){
		return err
	}

	valid =  p.GetByCodeValue(data.Code_value) 
	if !valid {
		return errors.New("el producto ya existe")
	}

	NuevoProducto := domain.Productos{
		Id:           len(p.ArrayProductos) + 1,
		Name:         data.Name,
		Quantity:     data.Quantity,
		Code_value:   data.Code_value,
		Is_published: data.Is_published,
		Expiration:   data.Expiration,
		Price:        data.Price,
	}

	p.ArrayProductos = append(p.ArrayProductos, NuevoProducto)

	return nil
}

// validar vacios
func validateEmptys(product *domain.Productos) (bool, error) {
	switch {
	case product.Name == "" || product.Code_value == "" || product.Expiration == "":
		return false, errors.New("los campos no pueden estar vacios")
	case product.Quantity <= 0 || product.Price <= 0:
		if product.Quantity <= 0 {
			return false, errors.New("la cantidad debe ser mayor que 0")
		}
		if product.Price <= 0 {
			return false, errors.New("la cantidad debe ser mayor que 0")
		}
	}
	return true, nil
} 

//validar que la fecha de vencimiento
func validateExpiration(product *domain.Productos) (bool, error) {
	dates := strings.Split(product.Expiration, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("fecha de caducidad no válida, debe estar en formato: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("fecha de caducidad no válida, deben ser números")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("fecha de caducidad no válida, la fecha debe estar entre el 1 y el 31/12/9999")
	}
	return true, nil
}

//buscar si el codigo ya existe
func (p ProductosAll) GetByCodeValue(code string) bool {
	for _, v := range p.ArrayProductos {
		if v.Code_value == code {
			return false
		}
	}
	return true
}