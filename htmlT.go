package main

// DO NOT FORGET TO EXECUTE:
// $ sqlite3 htmlT.db
// sqlite> CREATE TABLE data (
// 	         number INTEGER PRIMARY KEY,
//           double INTEGER,
// 	         square INTEGER );
//
// Run with:
//   $ go run htmlT.go htmlT.db fixtures/html.gohtml

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Entry struct {
	Number int
	Double int
	Square int
}

var (
	DATA  []Entry
	tFile string
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Host: %s Path: %s\n", r.Host, r.URL.Path)
	myT := template.Must(template.ParseGlob(tFile))
	myT.ExecuteTemplate(w, tFile, DATA)
	myT.Execute(os.Stdout, DATA)
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need Database File + Template File!")
		return
	}

	database := arguments[1]
	tFile = arguments[2] // use: fixtures/html.gohtml

	db, err := sql.Open("sqlite3", database)
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Emptying database table.")
	_, err = db.Exec("DELETE FROM data;")
	if err != nil {
		fmt.Println(nil)
		return
	}

	fmt.Println("Populating", database)
	stmt, _ := db.Prepare("INSERT INTO data(number, double, square) values(?,?,?);")
	for i := 20; i < 50; i++ {
		_, _ = stmt.Exec(i, 2*i, i*i)
		fmt.Println("inserted", i)
	}

	rows, err := db.Query("SELECT * FROM data;")
	if err != nil {
		fmt.Println(nil)
		return
	}

	var (
		n, d, s int
	)

	fmt.Println("rows preparation")
	for rows.Next() {
		err = rows.Scan(&n, &d, &s)
		temp := Entry{Number: n, Double: d, Square: s}
		fmt.Println("row:", n, d, s)
		DATA = append(DATA, temp)
	}

	fmt.Println("running server ...")

	http.HandleFunc("/", myHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

}
