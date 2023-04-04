package database

import (
	"database/sql"
	"fmt"

	"github.com/marcelofelixsalgado/financial-commons/pkg/commons/logger"
	"github.com/marcelofelixsalgado/financial-commons/settings"

	_ "github.com/go-sql-driver/mysql"
)

func NewConnection() *sql.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		settings.Config.DatabaseConnectionUser,
		settings.Config.DatabaseConnectionPassword,
		settings.Config.DatabaseConnectionServerAddress,
		settings.Config.DatabaseConnectionServerPort,
		settings.Config.DatabaseName)

	fmt.Printf("connectionString: %s", connectionString)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Erro ao conectar no banco")
		logger.GetLogger().Fatalf("Error trying to connect to database: %v", err)
	}

	db.SetConnMaxLifetime(0)
	db.SetMaxIdleConns(3)
	db.SetMaxOpenConns(3)

	// Checks if connection is open
	if err = db.Ping(); err != nil {
		db.Close()
		fmt.Println("Erro ao checar no banco")
		logger.GetLogger().Fatalf("Error trying to check the database connection: %v", err)
	}

	return db
}
