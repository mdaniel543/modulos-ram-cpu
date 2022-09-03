package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type ram struct {
	total      int `json:"total"`
	used       int `json:"used"`
	free       int `json:"free"`
	percentage int `json:"percentage"`
}

var conn = MySQLConn()

func MySQLConn() *sql.DB {
	db, err := sql.Open("mysql", "admin:7P4,;C<8Io^jG&p&@tcp(35.202.232.209)/modules")
	if err != nil {
		fmt.Println(err)
		fmt.Println("Error en la conexión a la base de datos")
	} else {
		fmt.Println("Connected to MySQL")
	}
	return db
}

func postRam(w http.ResponseWriter, r *http.Request) {
	var ram ram
	_ = json.NewDecoder(r.Body).Decode(&ram)
	stmt, err := conn.Prepare("INSERT INTO ram(total, used, free, percentage) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(ram.total, ram.used, ram.free, ram.percentage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ram)
}

func getRam(w http.ResponseWriter, r *http.Request) {
	var ram ram
	var err = conn.QueryRow("SELECT total, used, free, percentage FROM ram ORDER BY id DESC LIMIT 1").Scan(&ram.total, &ram.used, &ram.free, &ram.percentage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ram)
}

func getRams(w http.ResponseWriter, r *http.Request) {
	var rams []ram
	rows, err := conn.Query("SELECT total, used, free, percentage FROM ram")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var ram ram
		err = rows.Scan(&ram.total, &ram.used, &ram.free, &ram.percentage)
		if err != nil {
			fmt.Println(err)
		}
		rams = append(rams, ram)
	}
	fmt.Println(rams)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/rams", getRams).Methods("GET")
	router.HandleFunc("/ram", getRam).Methods("GET")
	router.HandleFunc("/ram", postRam).Methods("POST")
	fmt.Println("Server on port", 8000)
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		fmt.Println(err)
	}
}
