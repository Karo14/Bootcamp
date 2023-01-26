package Productos

import "clase2_web/internal/domain"

//STRUCTURA LLAMANDO LA INTERFACE DE REPOSITORY
type services struct{
	InterfaceRepository MetodosRepository
}

type MetodosProductosService interface{
	GetAllService() []domain.Productos
	GetByIdService(id int) (domain.Productos,error)
	GetMayorService(min float64) []domain.Productos
	CrearNuevoProductoService(data domain.Productos) error
	GetByCodeValueService(code string) bool
}

//CONSTRUCTOR
func NuevoServicio(Repository MetodosRepository) *services {
	return &services{InterfaceRepository:Repository}
}

//mostrar todo los productos
func (p *services)GetAllService() []domain.Productos{
	return p.InterfaceRepository.GetAllRepository()
}

//buscar un productos por un id
func (p services) GetByIdService(id int) (domain.Productos,error) {
	return p.InterfaceRepository.GetByIdRepository(id)
}

//buscar por valor mayor que
func (p services) GetMayorService(min float64) []domain.Productos {
	return p.InterfaceRepository.GetMayorRepository(min)
}

//crear nuevo producto
func (p services) CrearNuevoProductoService(data domain.Productos) error  {
	return p.InterfaceRepository.CrearNuevoProductoRepository(data)
}

//buscar si el codigo ya existe
func (p services) GetByCodeValueService(code string) bool {
	return p.InterfaceRepository.GetByCodeValue(code)
}


