package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
	db  *sql.DB
)

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

	db, err = sql.Open(dbName, conStr)
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
