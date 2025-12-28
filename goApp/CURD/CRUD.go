package main

// go mod init MyApp => go 모듈 자동 생성
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" //go get github.com/go-sql-driver/mysql => MySql 드라이버 설치
)

func main() {
	dsn := "root:It@stohard03@tcp(127.0.0.1:3306)/testdb"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var id int
	var name string

	CreateUser(db, 24, "harry")

	DeleteUser(db, 24)

	UpdateUser(db, 11, "lary")

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

func CreateUser(db *sql.DB, ID int, Name string) string {

	result, err := db.Exec("INSERT INTO test1 VALUES (?, ?)", ID, Name)
	if err != nil {
		return "insert fail"
	}
	n, err := result.RowsAffected()
	if n == 1 {
		return "1 row inserted"
	}
	return ""
}

func DeleteUser(db *sql.DB, ID int) string {

	result, err := db.Exec("DELETE FROM test1 WHERE id = ?", ID)
	if err != nil {
		return "delete fail"
	}
	n, err := result.RowsAffected()
	if n == 1 {
		return "deleted"
	}
	return ""

}

func UpdateUser(db *sql.DB, ID int, Name string) string {

	result, err := db.Exec("UPDATE test1 SET name = ?  WHERE id = ?", Name, ID)
	if err != nil {
		return "update fail"
	}
	n, err := result.RowsAffected()
	if n == 1 {
		return "updated"
	}
	return ""
}
