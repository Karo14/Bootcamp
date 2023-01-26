package products

import (
	"database/sql"
	//"database/sql/driver"
	"github.com/go-sql-driver/mysql"
)

type MySQLRepository struct{
	Repository
	Database *sql.DB
}

//Constructor
func NewRepositorySQL(db *sql.DB) Repository {
	return &MySQLRepository{Database: db}
}

func (repository *MySQLRepository) Get(id int) (product *Product, err error){
	query :=  `
		SELECT id, name, type, count, price 
		FROM products 
		WHERE id = ?
		` 
	row := repository.Database.QueryRow(query, id)
	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			err = ErrNotFound
		default:
			err = ErrInternal
		}
		return
	}
	s := Product{}
	err = row.Scan(
		&s.ID,
		&s.Name,
		&s.Type,
		&s.Count,
		&s.Price,
	)

	product = &s
	if  err != nil {
		err = ErrInternal
		return
	}

	return
}

func (repository*MySQLRepository) Store(product *Product) (err error) {
	statement, err :=  repository.Database.Prepare(`
		INSERT INTO products (
			name, type, count, price)
		VALUES (?,?,?,?); 
	`)
	if err != nil{
		err = ErrInternal
		return 
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Type,
		product.Count,
		product.Price,
	)

	if err != nil{
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			err = ErrInternal
			return 
		}

		switch driverErr.Number{
		case 1862:
			err = ErrDuplicated
		default:
			err = ErrInternal
		}
	}

	id, err := result.LastInsertId()
	product.ID = int(id)

	return 
}

func (repository *MySQLRepository) Update(product *Product) (err error){
	statement, err := repository.Database.Prepare(`UPDATE products SET name = ?, type= ?, count= ?, price= ? WHERE id =?; `)
	if err != nil{
		err = ErrInternal
		return 
	}
	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Type,
		product.Count,
		product.Price,
		product.ID,
	)

	_ , err = result.RowsAffected()

	if err != nil{
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			err = ErrInternal
			return 
		}

		switch driverErr.Number{
		case 1862:
			err = ErrDuplicated
		default:
			err = ErrInternal
		}
	}
	return
}

func (repository *MySQLRepository) Delete(id int) (err error){
	query :=  `DELETE FROM products WHERE id = ?` 
	
	row := repository.Database.QueryRow(query, id)
	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			err = ErrNotFound
		default:
			err = ErrInternal
		}
		return
	}
	return
}
