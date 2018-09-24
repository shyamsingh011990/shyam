package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	  "net/http"
    "log"
    "github.com/gorilla/mux"
    "encoding/json"
)

var db *sql.DB
var stmt *sql.DB



type test_struct  struct {
	//Token string
	 Token string 
}

func setupResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:8000/insert")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    (*w).Header().Set("Content-Type", "application/json")
}

func handler(w http.ResponseWriter, r *http.Request){
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	//w.WriteHeader(http.StatusOK)
//msg :=r.body.String()
//fmt.Println(r.Body)
	//msg:=token
	
	
	decoder := json.NewDecoder(r.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.Token)
	owner := t.Token
	//name := data.Name

	// Prepare statement for inserting data
	stmt , err := db.Prepare("insert into tokens (token) values(?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmt.Exec(owner)

		if err != nil {
			fmt.Print(err.Error())
		}
	 
}







func main() {


	

	//db, err := sql.Open("mysql", "user:password@/database")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())  // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	
r := mux.NewRouter()
	//r.HandleFunc("/", handler)
	r.HandleFunc("/insert", handler).Methods("POST","OPTIONS")

log.Fatal(http.ListenAndServe("localhost:8000", r))
}
