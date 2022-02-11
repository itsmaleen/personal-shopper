package graph

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/itsmaleen/personal-shopper/backend/graph/model"
)

func createSchema(db *pg.DB) error {
	for _, models := range []interface{}{(*model.ImageData)(nil), (*model.Tag)(nil)} {
		if err := db.Model(models).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			panic(err)
		}
	}

	return nil
}

func Connect() *pg.DB {
	DB_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	DB_PORT := os.Getenv("POSTGRES_PORT")
	DB_NAME := os.Getenv("POSTGRES_DB")
	DB_HOST := os.Getenv("POSTGRES_HOST")
	DB_USER := os.Getenv("POSTGRES_USER")

	connStr := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)

	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	if schemaErr := createSchema(db); schemaErr != nil {
		panic(schemaErr)
	}

	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		panic("PostgreSQL is down")
	}

	return db
}
