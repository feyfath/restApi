package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

type Users struct {
	ID    int
	Title string
	Desc  string
	Text  string
}
type Pictures struct {
	ID    int
	Title string
	Url   string
}

func checkErr(err error) {
	if err != nil {
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home page!")
}

func cours(w http.ResponseWriter, r *http.Request) {
	var myUser Users
	coursRows, err := db.Query("select * from cours")
	checkErr(err)
	defer coursRows.Close()
	for coursRows.Next() {
		coursRows.Scan(&myUser.ID, &myUser.Title, &myUser.Desc, &myUser.Text)
		json.NewEncoder(w).Encode(myUser)
	}
}

func pictures(w http.ResponseWriter, r *http.Request) {
	var myPicture []*Pictures
	coursRows, err := db.Query("select * from pictures")
	checkErr(err)
	defer coursRows.Close()
	for coursRows.Next() { // this stops when there are no more rows
		c := new(Pictures)                             // initialize a new instance
		err := coursRows.Scan(&c.ID, &c.Title, &c.Url) // scan contents of the current row into the instance
		checkErr(err)
		myPicture = append(myPicture, c) // add each instance to the slice
	}
	if err := json.NewEncoder(w).Encode(myPicture); err != nil {
		log.Println(err)
	}
}

func handler() {
	fmt.Printf("hello handler!")
	http.HandleFunc("/", home)
	http.HandleFunc("/cours", cours)
	http.HandleFunc("/pics", pictures)
	http.ListenAndServe(":8082", nil)
}

func main() {
	db, err = sql.Open("mysql", "root:8520@tcp(127.0.0.1:3306)/world")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("successfully connected to database.")
	}
	defer db.Close()

	handler()
}
