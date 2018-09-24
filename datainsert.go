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



type myData struct {
	Token string
	//date  string
}

func handler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
//msg :=r.body.String()
//fmt.Println(r.Body.String())
	//msg:=token
	
	
	decoder := json.NewDecoder(r.Body)
	var data myData
	err := decoder.Decode(&data)
	if err != nil {
		panic(err)
	}

	owner := data.Token
	//name := data.Name

	// Prepare statement for inserting data
	st , err := db.Prepare("insert into tokens (token) values(?)") // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	_, err = stmt.Exec(owner)

		if err != nil {
			fmt.Print(err.Error())
		}
		defer st.Close()
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
	r.HandleFunc("/insert", handler).Methods("POST")

log.Fatal(http.ListenAndServe("localhost:8000", r))
}
