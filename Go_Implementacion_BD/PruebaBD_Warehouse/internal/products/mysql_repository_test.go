package products

import (
	"PruebaBD_Warehouse/internal/domain"
	"database/sql"
	"testing"

	"time"

	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func init(){
	databaseConfig := mysql.Config{
		User: "root",
		Passwd: "",
		Addr: "localhost:3306",
		DBName: "my_db",
		ParseTime: true,
	}
	txdb.Register("txdb","mysql",databaseConfig.FormatDSN())
}

func TestGet(t *testing.T) {
	//arrange
	//Open Database connection
	database, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer database.Close()

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
	database, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer database.Close()
	
	rp := NewRepositorySQL(database)
	
	// act
	tm,_ := time.Parse("2006-01-02","2023-01-26")
	s := domain.Product{
		Name : "Zapatos",
		Quantity : 5,
		Code_value : "05",
		Is_published: "NO",
		Expiration: tm,
		Price : 109000,
		Id_warehouse: 1,
	}

	// assert
	err = rp.Store(&s)
	assert.NoError(t, err)
	t.Logf("products: %+v", s)
}

func TestUpdate(t *testing.T) {
	//arrange
	//Open Database connection
	database, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer database.Close()
	
	rp := NewRepositorySQL(database)

	// act
	tm,_ := time.Parse("2006-01-02","2023-01-26")
	p := domain.Product{
		ID: 2,
		Name : "Esqueleto",
		Quantity : 5,
		Code_value : "05",
		Is_published: "NO",
		Expiration: tm,
		Price : 14000,
		Id_warehouse: 1,
	}
	
	// assert
	err = rp.Update(&p)
	assert.NoError(t, err)
	t.Logf("products: %+v", p)
}

func TestDelete(t *testing.T) {
	//arrange
	//Open Database connection
	database, err := sql.Open("txdb", "identifier")
	assert.NoError(t, err)
	defer database.Close()
	
	rp := NewRepositorySQL(database)

	// act
	err = rp.Delete(200)

	// assert
	assert.NoError(t, err)
}