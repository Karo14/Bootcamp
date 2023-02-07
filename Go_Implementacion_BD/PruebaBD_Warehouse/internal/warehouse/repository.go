package warehouse

import (
	"PruebaBD_Warehouse/internal/domain"
	"database/sql"
	"errors"
)

const (
	GetReport 	   = "SELECT w.id, w.name, w.address, COUNT(p.warehouse_id) as product_count FROM warehouse as w " +
					 "LEFT JOIN products as p ON w.id = p.warehouse_id " +
					 "GROUP BY w.id"
	GetReportQuery = "SELECT w.id, w.name, w.address, COUNT(p.warehouse_id) as product_count FROM warehouse as w " +
					 "LEFT JOIN products as p ON w.id = p.warehouse_id " +
					 "GROUP BY w.id HAVING w.id = ?;"
)

var (
	ErrInternal = errors.New("internal error")
)

type Repository interface {
	// read
	GetProductReport(id int) ([]domain.WarehouseReport, error)
}

func NewRepositorySQL(db *sql.DB) Repository {
	return &repositorySQL{db: db}
}

type repositorySQL struct {
	db *sql.DB
}

func (rp *repositorySQL) GetProductReport(id int) ([]domain.WarehouseReport, error) {
	var rows *sql.Rows; var err error
	if id > 0 {
		rows, err = rp.db.Query(GetReportQuery, id)
	} else {
		rows, err = rp.db.Query(GetReport)
	}
	if err != nil {
		return nil, ErrInternal
	}

	var reports []domain.WarehouseReport
	for rows.Next() {
		var report domain.WarehouseReport
		err := rows.Scan(&report.ID, report.Name, report.Address, report.ProductCount)
		if err != nil {
			return nil, ErrInternal
		}

		reports = append(reports, report)
	}

	return reports, nil
}