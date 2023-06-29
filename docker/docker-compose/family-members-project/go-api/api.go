package main

import "log"
import "net/http"
import "encoding/json"
//import "strconv"
import "github.com/rs/cors"
import "github.com/gorilla/mux"
import "database/sql"
// import "fmt"
import _ "github.com/go-sql-driver/mysql"

type Member struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    string `json:"age"`
	Gender string `json:"gender"`
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	log.Println("getMembers function called")

	// Create an array to return as json response
	var members []Member

	// Open a connection to mysql
	db, err := sql.Open("mysql", "api:apipass@tcp(192.168.0.2:3306)/api")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Send query to mysql via connection
	rows, err := db.Query("SELECT * FROM family")
	if err != nil {
		log.Fatal(err)
	}
	// For loop over query contents and print
	for rows.Next() {
		// Create Member struct
		var m Member
		// Scan query data into struct
		err = rows.Scan(&m.ID, &m.Name, &m.Age, &m.Gender)
		if err != nil {
			log.Fatal(err)
		}
		// Append struct to Members array
		members = append(members, m)
	}
	json.NewEncoder(w).Encode(members)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/members", getMembers).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://192.168.0.2:3000", "http://127.0.0.1:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions, http.MethodHead},
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)
	log.Fatal(http.ListenAndServe(":8000", handler))
}
