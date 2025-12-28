package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:It@stohard03@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	var name string

	result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", 11, "jack")
	if err != nil {
		log.Fatal(err)
	}
	n, err := result.RowsAffected()
	if n == 1 {
		fmt.Println("1 row inserted")
	}

	rows, err := db.Query("SELECT id, name FROM test1 WHERE id >= ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

}
