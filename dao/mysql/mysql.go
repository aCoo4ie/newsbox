package mysql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// Global database connection
var db *sqlx.DB

// Init initializes the database connection and stores it in the package-level variable
func Init() error {
	host := viper.GetString("mysql.host")
	port := viper.GetInt("mysql.port")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	dbname := viper.GetString("mysql.dbname")

	dbs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4", user, password, host, port, dbname)

	var err error
	db, err = sqlx.Connect("mysql", dbs)
	if err != nil {
		fmt.Printf("Failed to init mysql: %v\n", err)
		return err
	}
	return nil
}

// GetDB returns the database connection
func GetDB() *sqlx.DB {
	return db
}

func Close() {
	_ = db.Close()
}
