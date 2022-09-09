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

type Process struct {
	Pid      int       `json:"pid"`
	Name     string    `json:"name"`
	User     int       `json:"user"`
	State    int       `json:"state"`
	Memory   float32   `json:"memory"`
	Children []Process `json:"children"`
}

type CPU struct {
	Usage int `json:"usage"`
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
	fmt.Println("Insertando RAM en la base de datos")
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
	fmt.Println("Ram insertada")
}

func deleteProcesses() {
	fmt.Println("Eliminando procesos de la base de datos...")
	stmt, err := conn.Prepare("DELETE FROM process")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Procesos eliminados")
}

func postProcesses(data string) {
	var processes []Process
	json.Unmarshal([]byte(data), &processes)
	deleteProcesses()
	fmt.Println("Insertando procesos en la base de datos ...")
	for _, process := range processes {
		stmt, err := conn.Prepare("INSERT INTO process(pid, name, user, state, memory) VALUES(?, ?, ?, ?, ?)")
		if err != nil {
			fmt.Println(err)
		}
		_, err = stmt.Exec(process.Pid, process.Name, process.User, process.State, (process.Memory/(1024.0*1024.0))*100.0)
		if err != nil {
			fmt.Println(err)
		}
		for _, child := range process.Children {
			stmt, err := conn.Prepare("INSERT INTO process(pid, name, user, state, memory, pid_padre) VALUES(?, ?, ?, ?, ?, ?)")
			if err != nil {
				fmt.Println(err)
			}
			_, err = stmt.Exec(child.Pid, child.Name, child.User, child.State, (child.Memory/(1024.0*1024.0))*100.0, process.Pid)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
	fmt.Println("Procesos insertados")
}

func postCPU(data string) {
	fmt.Println("Insertando CPU en la base de datos")
	fmt.Println(data)
	var cpu CPU
	json.Unmarshal([]byte(data), &cpu)
	fmt.Println(cpu)

	stmt, err := conn.Prepare("INSERT INTO cpu(percentage) VALUES(?)")
	if err != nil {
		fmt.Println(err)
	}
	_, err = stmt.Exec(100 - cpu.Usage)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("CPU insertada")

}

func main() {
	for {
		fmt.Println("Obteniendo datos ...")
		cmd := exec.Command("sh", "-c", "cat /proc/ram_201709450")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		postRam(string(out[:]))
		// ------------------------------------------------------------------
		proccess := exec.Command("sh", "-c", "cat /proc/cpu_201709450")
		out, err = proccess.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		postProcesses(string(out[:]))
		// ------------------------------------------------------------------
		cpu := exec.Command("sh", "-c", "echo \"{ usage:$(vmstat 1 2|tail -1|awk '{print $15}') }\"")
		out, err = cpu.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		postCPU(string(out[:]))
		// ------------------------------------------------------------------
		time.Sleep(6 * time.Second)
	}
}
