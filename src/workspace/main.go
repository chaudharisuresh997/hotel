package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func (dao Dao) GetMySqlConnection() {
	db, err := sql.Open("mysql", "newuser:newpassword@tcp(127.0.0.1:3306)/test_db")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	dao.mc = db
	return
}

type Dao struct {
	mc *sql.DB
}

var dao Dao

func init() {
	dao.GetMySqlConnection()
}
func stockHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	//fmt.Printf("entered %v", vars)
	json.NewEncoder(w).Encode(vars)

}
func main() {
	r := mux.NewRouter()
	//example 1: /search?query=sample?stock="infy"
	r.HandleFunc("/search", stockHandler).
		Queries(
			"query", "{query}",
			"stock", "{stock}",
		).
		Methods("GET")
	//	r.HandleFunc("/user", UserByValueHandler).Methods("GET")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:9000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
