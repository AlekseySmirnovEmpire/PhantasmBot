package db

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"io/ioutil"
)

type dbConnect struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}

const dbName string = "postgres"

var (
	con *dbConnect
	db  *sqlx.DB
)

func CloseDB() {
	db.Close()
}

func InitDB() error {
	if con != nil {
		return nil
	}

	fmt.Println("Connecting DB ....")

	file, err := ioutil.ReadFile("./dbConfig.json")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	con = new(dbConnect)
	err = json.Unmarshal(file, con)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	conStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		con.Host, con.Port, con.User, con.Password, con.DbName)

	db, err = sqlx.Open(dbName, conStr)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	if err = db.Ping(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Connecting DB SUCCESS!")
	return nil
}

func Update(query *string) error {
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

func getZero[T any]() T {
	var result T
	return result
}
