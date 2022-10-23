package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const dbName string = "postgres"

var (
	db *sqlx.DB
)

func CloseDB() {
	_ = db.Close()
}

func InitDB() error {
	if db != nil {
		return nil
	}

	log.Println("Connecting DB ....")

	conStr, exist := os.LookupEnv("DB_ConnectionString")
	if !exist {
		log.Println("There is no connection string in .env file!")
	}

	var err error
	db, err = sqlx.Open(dbName, conStr)
	if err != nil {
		log.Println("SQL connection not opened!")
		return err
	}

	if err = db.Ping(); err != nil {
		log.Printf("PING DB end with error: \"%s\"", err.Error())
		return err
	}

	log.Println("Connecting DB SUCCESS!")
	return nil
}

func InsertOrUpdate(query *string) error {
	_, err := db.Exec(*query)
	if err == nil {
		return err
	}
	return nil
}

func Select[T comparable](sql *string) ([]T, error) {
	objects := make([]T, 0)
	rows, err := db.Queryx(*sql)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var o T
		err = rows.StructScan(&o)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		objects = append(objects, o)
	}

	return objects, nil
}

//func getZero[T any]() T {
//	var result T
//	return result
//}
