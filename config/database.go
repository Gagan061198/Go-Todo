package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {

	user := "root"
	pass := "root"

	credentials := fmt.Sprintf("%s:%s@/gotodo?charset=utf8&parseTime=True", user, pass)

	database, err := sql.Open("mysql", credentials)

	if err != nil {

		log.Fatal(err)
	} else {
		fmt.Println("Database Connection Successful")
	}

	_, err = database.Exec(`CREATE DATABASE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`USE gotodo`)

	if err != nil {
		fmt.Println(err)
	}

	_, err = database.Exec(`
		CREATE TABLE todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)

	if err != nil {
		fmt.Println(err)
	}

	return database
}
