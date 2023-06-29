package main

import "log"
import "net/http"
import "encoding/json"
//import "strconv"
import "github.com/rs/cors"
import "github.com/gorilla/mux"

type Member struct {
    ID      int     `json:"id"`
    Name   string  `json:"name"`
    Age  string  `json:"age"`
    Gender    string  `json:"gender"`
}

var members []Member

func getMembers(w http.ResponseWriter, r *http.Request) {
	log.Println("getMembers function called")
    json.NewEncoder(w).Encode(members)
}

func main() {
    r := mux.NewRouter()

    members = append(members,
        Member{ID: 1, Name: "Andre Kalaj", Age: "72", Gender: "Male"},
        Member{ID: 2, Name: "Teresa Kalaj", Age: "66", Gender: "Female"},
        Member{ID: 3, Name: "Chris Kalaj", Age: "41", Gender: "Male"},
        Member{ID: 4, Name: "Alex Kalaj", Age: "33", Gender: "Male"},
        Member{ID: 5, Name: "Lily Kalaj", Age: "31", Gender: "Female"})

    r.HandleFunc("/members", getMembers).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://192.168.0.2:3000", "http://127.0.0.1:3000"}, 
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions, http.MethodHead}, 
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(r)
    log.Fatal(http.ListenAndServe(":8000", handler))
}
