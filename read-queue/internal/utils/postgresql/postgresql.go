package postgresql

import (
	"fmt"
	"readq/internal/config"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var dbx *DBX
var once sync.Once

// DBX wrapper class for sqlx.DB
type DBX struct {
	sqlx.DB
}

// InitDatabase Function for init database
func InitDatabase(connectionString string) {
	once.Do(func() {
		dbs, err := sqlx.Connect("postgres", connectionString)
		if err != nil {
			fmt.Println("Error connecting to database: ", err)
			panic(err)
		}

		dbx = &DBX{
			DB: *dbs,
		}

		config := config.GetConfig()
		if config.PostgresDB.MaxOpenConn > 0 {
			dbx.SetMaxOpenConns(config.PostgresDB.MaxOpenConn) // The default is 0 (unlimited)
		}

		if config.PostgresDB.MaxIdleConn > 0 {
			dbx.SetMaxIdleConns(config.PostgresDB.MaxIdleConn) // defaultMaxIdleConns = 2
		}

		if config.PostgresDB.ConnMaxLifeTimeTTL != nil {
			dbx.SetConnMaxLifetime(*config.PostgresDB.ConnMaxLifeTimeTTL) // 0, connections are reused forever.
		}
	})
}

func Ping() error {
	if dbx == nil {
		return nil
	}

	err := dbx.Ping()
	if err != nil {
		return err
	}

	return nil
}

// UnInitDatabase cleanup database
func UnInitDatabase() {
	if dbx != nil {
		dbx.DB.Close()
	}
}

// Open database
func Open() (*DBX, error) {

	return dbx, nil
}

// Close Override Close function
func (dbx *DBX) Close() {
	//do nothing
}
