package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Ram struct {
	Total      int `json:"total"`
	Free       int `json:"free"`
	Used       int `json:"used"`
	Percentage int `json:"percentage"`
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

func postRam(data string) {
	fmt.Println("Insertando datos en la base de datos")
	fmt.Println(data)
	var ram Ram
	json.Unmarshal([]byte(data), &ram)
	fmt.Println(ram)

	stmt, err := conn.Prepare("INSERT INTO ram(total, used, free, percentage) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(ram.Total, ram.Used, ram.Free, ram.Percentage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Datos insertados")
}

func main() {
	for {
		fmt.Println("Obteniendo datos ...")
		cmd := exec.Command("sh", "-c", "cat /proc/ram_201709450")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		output := string(out[:])
		fmt.Println(output)
		postRam(output)
		time.Sleep(6 * time.Second)
	}
}
