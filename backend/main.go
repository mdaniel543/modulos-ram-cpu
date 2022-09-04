package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
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

func postRam(data string) {
	fmt.Println("Insertando datos en la base de datos")
	fmt.Println(data)
	var ram ram
	json.Unmarshal([]byte(data), &ram)
	fmt.Println(ram.free, ram.percentage, ram.total, ram.used)

	/*stmt, err := conn.Prepare("INSERT INTO ram(total, used, free, percentage) VALUES(?, ?, ?, ?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(ram.total, ram.used, ram.free, ram.percentage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Datos insertados")*/
}

func main() {
	fmt.Println("Datos obtenidos de la memoria RAM")

	cmd := exec.Command("sh", "-c", "cat /proc/ram_201709450")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	output := string(out[:])
	fmt.Println(output)
	postRam(output)
}
