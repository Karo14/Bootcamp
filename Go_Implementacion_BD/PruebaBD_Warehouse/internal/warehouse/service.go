package warehouse

import "PruebaBD_Warehouse/internal/domain"


type Service interface {
	GetReport(id int) ([]domain.WarehouseReport, error)
}

func NewService(rp Repository) Service {
	return &service{rp: rp}
}

type service struct {
	rp Repository
}
// read
func (sv *service) GetReport(id int) ([]domain.WarehouseReport, error) {
	return sv.rp.GetProductReport(id)
}