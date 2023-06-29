package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Member struct {
    ID      int     `json:"id"`
    Name   string  `json:"name"`
    Age  string  `json:"age"`
    Gender    string  `json:"gender"`
}

func main() {
	db, err := sql.Open("mysql", "api:apipass@tcp(192.168.0.2:3306)/api")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM family")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var m Member
		err = rows.Scan(&m.ID,&m.Name,&m.Age,&m.Gender)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
	}

}
