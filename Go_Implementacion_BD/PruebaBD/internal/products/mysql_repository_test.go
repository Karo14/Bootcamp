package products

import (
	"database/sql"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	//arrange
	//Open Database connection
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "storage",
	}
	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	} 
	
	rp := NewRepositorySQL(database)

	// act
	products, err := rp.Get(1)

	// assert
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
	t.Logf("products: %+v", products)
}

func TestStore(t *testing.T) {
	//arrange
	//Open Database connection
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "storage",
	}
	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	} 
	
	rp := NewRepositorySQL(database)

	// act
	s := Product{
		Name : "Zapatos",
		Type : "Tenis",
		Count : 1,
		Price : 109000,
	}

	// assert
	err = rp.Store(&s)
	assert.NoError(t, err)
	t.Logf("products: %+v", s)
}

func TestUpdate(t *testing.T) {
	//arrange
	//Open Database connection
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "storage",
	}
	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	} 
	
	rp := NewRepositorySQL(database)

	// act
	p := Product{
		ID: 2,
		Name : "Esqueleto",
		Type : "Sin mangas",
		Count : 1,
		Price : 14000,
	}

	// assert
	err = rp.Update(&p)
	assert.NoError(t, err)
	t.Logf("products: %+v", p)
}

func TestDelete(t *testing.T) {
	//arrange
	//Open Database connection
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "storage",
	}
	database, err := sql.Open("mysql", databaseConfig.FormatDSN())
	if err != nil {
		panic(err)
	} 
	
	rp := NewRepositorySQL(database)

	// act
	err = rp.Delete(4)

	// assert
	assert.NoError(t, err)
}